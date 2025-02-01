package route

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/config"
	"github.com/sahidhossen/synmail/src/email"
	"github.com/sahidhossen/synmail/src/handler"
	"github.com/sahidhossen/synmail/src/middleware"
	"github.com/sahidhossen/synmail/src/migrations"
	"github.com/sahidhossen/synmail/src/services"
)

func Router(r *gin.RouterGroup, ctx context.Context, cfg *config.Config) {
	db := config.ConnectDB(cfg.DatabaseConnectionString)

	migrations.Migrate(db)

	service := services.SynMailServices{DB: db}
	emailService, err := email.NewEmailService(email.SMTP, cfg)

	if err != nil {
		log.Err(err).Msg("Email service error!")
	}

	middleware := middleware.CreateMiddleware(cfg)

	ginHandler := handler.CreateHandler(cfg, &emailService, &service)
	r.GET("/ping", ginHandler.Ping)

	userApi := r.Group("/user")
	userApi.POST("/register", ginHandler.RegisterUser)
	userApi.POST("/login", ginHandler.Login)
	userApi.GET("/me", middleware.AuthMiddleware, ginHandler.UserInfo)
	userApi.GET("/:id", ginHandler.GetUser)

	campaign := r.Group("/campaign")
	campaign.POST("/create", middleware.AuthMiddleware, ginHandler.CreateCampaign)
	campaign.GET("/all", middleware.AuthMiddleware, ginHandler.GetCampaigns)
	campaign.PATCH("/:id", middleware.AuthMiddleware, ginHandler.UpdateCampaign)
	campaign.DELETE("/:id", middleware.AuthMiddleware, ginHandler.DeleteCampaign)
	campaign.GET("/:id", middleware.AuthMiddleware, ginHandler.GetCampaign)

	subscriber := r.Group("/subscriber")
	subscriber.POST("/create", middleware.AuthMiddleware, ginHandler.CreateSubscribe)
	subscriber.POST("/import", middleware.AuthMiddleware, ginHandler.ImportSubscriber)
	subscriber.GET("/all", middleware.AuthMiddleware, ginHandler.GetSubscribers)
	subscriber.PATCH("/:id", middleware.AuthMiddleware, ginHandler.UpdateSubscribe)
	subscriber.DELETE("/:id", middleware.AuthMiddleware, ginHandler.DeleteSubscribe)
	subscriber.GET("/:id", middleware.AuthMiddleware, ginHandler.GetSubscribe)

	topics := r.Group("/topics")
	topics.POST("/create", middleware.AuthMiddleware, ginHandler.CreateTopics)
	topics.GET("/all", middleware.AuthMiddleware, ginHandler.GetSubscribeTopices)
	topics.PATCH("/:id", middleware.AuthMiddleware, ginHandler.UpdateTopic)
	topics.DELETE("/:id", middleware.AuthMiddleware, ginHandler.DeleteTopic)
	topics.GET("/:id", middleware.AuthMiddleware, ginHandler.GetSubscribeTopic)

	topicMap := r.Group("/topic_map")
	topicMap.POST("/create", middleware.AuthMiddleware, ginHandler.CreateTopicMap)
	topicMap.PATCH("/:id", middleware.AuthMiddleware, ginHandler.UpdateTopicMap)
	topicMap.DELETE("/:id", middleware.AuthMiddleware, ginHandler.DeleteTopicMap)
	topicMap.GET("/:id", middleware.AuthMiddleware, ginHandler.GetSubscribeTopicMap)

	template := r.Group("/template")
	template.POST("/create", middleware.AuthMiddleware, ginHandler.CreateTemplate)
	template.GET("/all", middleware.AuthMiddleware, ginHandler.GetTemplates)
	template.PATCH("/:id", middleware.AuthMiddleware, ginHandler.UpdateTemplate)
	template.DELETE("/:id", middleware.AuthMiddleware, ginHandler.DeleteTemplate)
	template.GET("/:id", middleware.AuthMiddleware, ginHandler.GetTemplate)

	tracker := r.Group("/tracker")
	tracker.POST("/create", middleware.AuthMiddleware, ginHandler.CreateTracker)
	tracker.GET("/all", middleware.AuthMiddleware, ginHandler.GetTrackers)
	tracker.PATCH("/:id", middleware.AuthMiddleware, ginHandler.UpdateTracker)
	tracker.DELETE("/:id", middleware.AuthMiddleware, ginHandler.DeleteTracker)
	tracker.GET("/:id", middleware.AuthMiddleware, ginHandler.GetTracker)

	unsubscriber := r.Group("/unsubscriber")
	unsubscriber.POST("/create", middleware.AuthMiddleware, ginHandler.CreateUnsubscribe)
	unsubscriber.PATCH("/:id", middleware.AuthMiddleware, ginHandler.UpdateUnsubscribe)
	unsubscriber.DELETE("/:id", middleware.AuthMiddleware, ginHandler.DeleteUnsubscribe)
	unsubscriber.GET("/:id", middleware.AuthMiddleware, ginHandler.GetUnSubscribe)
}
