package main

import (
	handlers "urlshortner/internal/controllers"
	"urlshortner/internal/logger"
	"urlshortner/internal/service"
	"urlshortner/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.NewLogger()

	storage := storage.NewInMemory()
	service := service.NewUrlShortner(logger, storage)

	handlers := handlers.NewUserHandler(service)

	router := gin.Default()

	router.POST("/create", handlers.CreateShortUrl)
	router.GET("/get", handlers.GetLongUrl)
	router.DELETE("/delete", handlers.DeleteUrl)

	//TODO: Add timeout wrapper/middleware
	router.Run()

}
