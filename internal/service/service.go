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

func (s *UrlShortener) Create(url string) (string, error) {
	s.logger.Info("creating url")
	short, err := s.storage.Create(url)
	if err != nil {
		s.logger.Error("creating url error")
		return "", err
	}
	return short, nil
}

func (s *UrlShortener) Get(url string) (string, error) {
	s.logger.Info("getting url")
	long, err := s.storage.Get(url)
	if err != nil {
		s.logger.Error("getting url error")
		return "", err
	}
	return long, nil
}

func (s *UrlShortener) Delete(url string) error {
	s.logger.Info("deleting url")
	err := s.storage.Delete(url)
	if err != nil {
		s.logger.Error("error deleting url")
		return err
	}
	return nil
}
