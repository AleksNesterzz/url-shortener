package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	handlers "urlshortner/internal/controllers"
	"urlshortner/internal/logger"
	auth "urlshortner/internal/service/auth"
	url "urlshortner/internal/service/url"
	"urlshortner/internal/token"

	"urlshortner/internal/storage"
	"urlshortner/internal/validator"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.New()
	tokenGen := &token.TokenGenerator{}
	storageURL := storage.NewInMemoryURLs()
	storageUser := storage.NewInMemoryUser()
	validator := validator.New()
	urlService := url.NewUrlShortener(logger, storageURL, validator)
	authService := auth.NewAuth(storageUser, logger, tokenGen)
	s := handlers.NewUserHandler(urlService)
	auth := handlers.NewAuthHandler(authService)

	router := gin.Default()
	//TODO: middleware auth
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
	router.POST("/logout", auth.Logout)
	router.POST("/create/", s.CreateShortUrl)
	router.GET("/get/:id", s.GetLongUrl)
	router.DELETE("/delete/:id", s.DeleteUrl)

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		logger.Info("HTTP server is starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start HTTP server: " + err.Error())
		}
	}()

	<-sig
	logger.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("HTTP server shutdown error: " + err.Error())
	} else {
		logger.Info("HTTP server stopped gracefully")
	}

}
