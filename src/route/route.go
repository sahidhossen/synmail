package route

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/email"
	"github.com/sahidhossen/synmail/src/handler"
	"github.com/sahidhossen/synmail/src/migrations"
	"github.com/sahidhossen/synmail/src/services"
)

func Router(r *gin.RouterGroup, ctx context.Context, cfg *config.Config) {
	db := config.ConnectDB(cfg.DatabaseConnectionString)

	migrations.Migrate(db)

	service := services.SynMailServices{DB: db}
	emailService, err := email.NewEmailService(email.SMTP, cfg)

	if err != nil {
		log.Fatal(err)
	}

	ginHandler := handler.CreateHandler(cfg, &emailService, &service)
	r.GET("/ping", ginHandler.Ping)

	userApi := r.Group("/user")
	userApi.POST("/register", ginHandler.RegisterUser)
	userApi.GET("/:id", ginHandler.GetUser)
}
