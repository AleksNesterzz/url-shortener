package service

import (
	"urlshortner/internal/logger"
	"urlshortner/internal/storage"
)

type UrlShortener struct {
	logger  logger.Logger
	storage storage.Storage
}

func NewUrlShortner(logger logger.Logger, storage storage.Storage) *UrlShortener {
	return &UrlShortener{
		logger:  logger,
		storage: storage,
	}
}

func Create(url string) (string, error) {
	return "", nil
}
