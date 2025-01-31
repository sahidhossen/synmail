package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mailgun/errors"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/config"
)

type MiddlewareHandler struct {
	*config.Config
}

type Auth struct {
	UserID      uint
	EmailID     string
	AccessToken jwt.Token
}

type UserClaim struct {
	UserID  uint   `json:"userId"`
	EmailID string `json:"emailId,omitempty"`
}

type TokenResponse struct {
	IssuedAt  *jwt.NumericDate `json:"issueAt"`
	ExpiresAt *jwt.NumericDate `json:"expiresAt"`
	Token     string           `json:"token"`
}

type SchClaims struct {
	Type string `json:"Name"`
	jwt.RegisteredClaims
	User UserClaim
}

type httpErrorResponse struct {
	Error string `json:"error"`
}

func CreateMiddleware(cfg *config.Config) *MiddlewareHandler {
	return &MiddlewareHandler{Config: cfg}
}

func GetAuth(c *gin.Context) *Auth {
	ptr := c.Value("auth").(*Auth)
	if ptr == nil {
		panic("cannot use auth in unauthoried context! did you use AuthMiddleware?")
	}
	return ptr
}

func (h *MiddlewareHandler) AuthMiddleware(c *gin.Context) {

	authHeader := strings.Split(c.GetHeader("Authorization"), " ")
	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization token is missing"})
	}
	auth, err := AuthenticateJwt(h.Secret, authHeader[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, &httpErrorResponse{Error: err.Error()})
		c.Abort()
		return
	}

	if auth.UserID == 0 {
		c.JSON(http.StatusUnauthorized, &httpErrorResponse{Error: "Authorization data invalid!"})
		c.Abort()
		return
	}

	c.Set("auth", auth)
	c.Next()
}

func AuthenticateJwt(secret string, jwtToken string) (*Auth, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &SchClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*SchClaims); ok && token.Valid {
		return &Auth{
			UserID:      claims.User.UserID,
			EmailID:     claims.User.EmailID,
			AccessToken: *token,
		}, nil
	} else {
		return nil, authErrorHandling(err)
	}
}

func CreateClaim(secret string, issuer string, user UserClaim) (TokenResponse, error) {
	signingKey := []byte(secret)
	// Create claims with multiple fields populated
	claims := SchClaims{
		"synMail",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
			Subject:   fmt.Sprintf("%d", user.UserID),
			ID:        "1",
			Audience:  []string{"internal"},
		},
		UserClaim{
			UserID:  user.UserID,
			EmailID: user.EmailID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)

	response := TokenResponse{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Token:     fmt.Sprintf("Bearer %s", ss),
	}
	return response, err
}

func authErrorHandling(err error) error {
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return errors.Wrap(err, "Invalid token format!")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		// Invalid signature
		log.Err(err).Msg("Invalid signature")
		return errors.Wrap(err, "Invalid signature!")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		log.Err(err).Msg("Token expored")
		return errors.Wrap(err, "Your token is expired!")
	} else {
		log.Err(err).Msg("Internal error when validate!")
		return errors.Wrap(err, "Internal error when validate!")
	}
}
