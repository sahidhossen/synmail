package route

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/email"
	"github.com/sahidhossen/synmail/src/handler"
)

func Router(r *gin.RouterGroup, ctx context.Context, cfg *config.Config) {
	emailService, err := email.NewEmailService(email.SMTP, cfg)
	if err != nil {
		log.Fatal(err)
	}
	ginHandler := handler.CreateHandler(cfg, &emailService)
	r.GET("/ping", ginHandler.Ping)
}
