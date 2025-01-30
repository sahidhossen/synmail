package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahidhossen/synmail/src/config"
	service "github.com/sahidhossen/synmail/src/service/email"
)

type GinHandler struct {
	*config.Config
	EmailClient *service.EmailClient
}

func CreateHandler(cfg *config.Config, emailClient *service.EmailClient) *GinHandler {
	return &GinHandler{Config: cfg, EmailClient: emailClient}
}

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ping from server!"})
}
