package service

import (
	"fmt"
	"urlshortner/internal/logger"
	"urlshortner/internal/storage"
	"urlshortner/internal/validator"
)

type UrlShortener struct {
	validator validator.Validator
	logger    logger.Logger
	storage   storage.Storage
}

func New(logger logger.Logger, storage storage.Storage, validator validator.Validator) *UrlShortener {
	return &UrlShortener{
		logger:    logger,
		storage:   storage,
		validator: validator,
	}
}

func (s *UrlShortener) Create(url string) (string, error) {
	s.logger.Info("creating short url")
	result := s.validator.Validate(url)
	if !result.IsValid {
		s.logger.Error("url is invalid")
		s.logger.Error(validator.StringArr(result.Errors).String())
		return "", fmt.Errorf("invalid url")
	}
	if len(result.Warnings) != 0 {
		s.logger.Warn("url have some warnings")
		s.logger.Warn(validator.StringArr(result.Warnings).String())
	}
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
