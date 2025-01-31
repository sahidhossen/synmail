package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/middleware"
	"github.com/sahidhossen/synmail/src/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *GinHandler) Login(c *gin.Context) {
	var reqBody *models.LoginRequest

	if err := c.Bind(&reqBody); err != nil {
		ResponseWithError(c, http.StatusBadRequest, "Password missing!")
		return
	}

	if reqBody.EmailID == "" && reqBody.UserName == "" {
		ResponseWithError(c, http.StatusBadRequest, "EmailID or userName required!")
		return
	}

	user, err := h.DBService.GetUserByEmailOrUserName(reqBody)
	if err != nil {
		log.Err(err).Msg("Internal query error")
		ResponseWithError(c, http.StatusInternalServerError, "Internal query error!")
		return
	}
	if user == nil {
		ResponseNotFound(c, "Authentication credential incorrect!")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
	if err != nil {
		ResponseWithError(c, http.StatusNotAcceptable, "Please type correct password")
		return
	}
	claimData := middleware.UserClaim{UserID: user.ID, EmailID: user.EmailID}

	token, err := middleware.CreateClaim(h.Config.Secret, "synUser", claimData)
	if err != nil {
		ResponseWithError(c, http.StatusBadRequest, "Internal error when creating claim!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": token})

}

func (h *GinHandler) RegisterUser(c *gin.Context) {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}

	// Hash the passward
	hash, hashError := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if hashError != nil {
		ResponseWithError(c, http.StatusBadRequest, "Failed to generate passward")
		return
	}
	user.Password = string(hash)
	err := h.DBService.CreateUser(user)
	if err != nil {
		log.Err(err).Msg("user create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "User created!")
}

func (h *GinHandler) GetUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	user, err := h.DBService.GetUserByID(uint(userID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "User query error!")
		return
	}
	if user == nil {
		ResponseNotFound(c, "User not found!")
		return
	}
	Response(c, http.StatusOK, "success", "User details", user)
}

func (h *GinHandler) UserInfo(c *gin.Context) {
	header := middleware.GetAuth(c)
	user, _ := h.DBService.GetUserByID(header.UserID)
	Response(c, http.StatusOK, "success", "User details", user)
}
