package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
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

	router.POST("/create/", handlers.CreateShortUrl)
	router.GET("/get/:id", handlers.GetLongUrl)
	router.DELETE("/delete", handlers.DeleteUrl)

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		time.Sleep(5 * time.Second)
		<-sig
		storage.Save()
	}()
	//TODO: Add timeout wrapper/middleware
	router.Run()

}
