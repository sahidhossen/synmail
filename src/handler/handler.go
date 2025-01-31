package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/email"
)

type GinHandler struct {
	*config.Config
	EmailService *email.EmailService
}

func CreateHandler(cfg *config.Config, emailClient *email.EmailService) *GinHandler {
	return &GinHandler{Config: cfg, EmailService: emailClient}
}

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ping from server!"})
}
