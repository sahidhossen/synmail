package route

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/handler"
	service "github.com/sahidhossen/synmail/src/service/email"
)

func Router(r *gin.RouterGroup, ctx context.Context, cfg *config.Config) {
	// repository, _ := service.CreateWithDefaultClient(ctx, *cfg)
	// emailRepository := service.NewEmailRepository(cfg, &service.EmailRepository{})
	emailClient := service.NewEmailClient(cfg, service.EmailClient{})
	ginHandler := handler.CreateHandler(cfg, emailClient)
	r.GET("/ping", ginHandler.Ping)
}
