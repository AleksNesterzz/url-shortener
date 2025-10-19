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
	logger := logger.New()
	storage := storage.New()
	service := service.New(logger, storage)

	handlers := handlers.New(service)

	router := gin.Default()
	//TODO: middleware auth + timeout context
	router.POST("/create/", handlers.CreateShortUrl)
	router.GET("/get/:id", handlers.GetLongUrl)
	router.DELETE("/delete/:id", handlers.DeleteUrl)

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
