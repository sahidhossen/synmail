package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/email"
	"github.com/sahidhossen/synmail/src/services"
)

type GinHandler struct {
	*config.Config
	EmailService email.EmailService
	DBService    *services.SynMailServices
}

func CreateHandler(cfg *config.Config, emailClient email.EmailService, service *services.SynMailServices) *GinHandler {
	return &GinHandler{Config: cfg, EmailService: emailClient, DBService: service}
}

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ping from server!"})
}
