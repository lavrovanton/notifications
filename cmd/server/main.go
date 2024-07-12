package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/lavrovanton/notifications/docs"
	"github.com/lavrovanton/notifications/internal/api"
	"github.com/lavrovanton/notifications/internal/config"
	"github.com/lavrovanton/notifications/internal/db"
	"github.com/lavrovanton/notifications/internal/rabbitmq"
	"github.com/lavrovanton/notifications/internal/rabbitmq/handler"
	"github.com/lavrovanton/notifications/internal/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title notifications
// @host localhost:9000
// @BasePath  /
func main() {
	cfg := config.Get()

	db, err := db.Get(cfg)
	if err != nil {
		log.Fatal("Start DB:", err)
	}

	notificationRepo := repository.NewNotificationRepository(db)
	notificationController := api.NewNotificationController(notificationRepo)

	router := gin.New()
	router.GET("/notifications", notificationController.Index)
	router.POST("/notifications", notificationController.Create)

	// Swagger
	docs.SwaggerInfo.Host = cfg.Host + ":" + cfg.Port
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// HTTP Server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("Server start: ", err)
		}
	}()

	// Rabbit MQ
	rmqUri := "amqp://" + cfg.RmqUser + ":" + cfg.RmqPassword + "@" + cfg.RmqHost + ":" + cfg.RmqPort + "/"
	consumer, err := rabbitmq.NewConsumer(
		rmqUri,
		"notify",
		handler.NewNotificationHandler(notificationRepo))
	if err != nil {
		log.Fatal("RabbitMQ start: ", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// HTTP Server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	}

	// Rabbit MQ
	if err = consumer.Shutdown(); err != nil {
		log.Println("RabbitMQ forced to shutdown:", err)
	}
}
