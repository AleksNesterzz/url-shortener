package main

import (
	handlers "urlshortner/internal/controllers"
	"urlshortner/internal/logger"
	"urlshortner/internal/service"
	"urlshortner/internal/storage"
)

func main() {
	logger := logger.NewLogger()

	storage := storage.NewInMemory()
	service := service.NewUrlShortner(logger, storage)

	_ = handlers.NewUserHandler(service)

}
