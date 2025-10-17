package logger

import (
	"log"
	"log/slog"
	"os"
)

type Logger interface {
	Info(string)
	Debug(string)
	Error(string)
}

type StdLogger struct {
	logger *slog.Logger
}

func NewLogger() *StdLogger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log := slog.New((slog.NewTextHandler(file, nil)))
	log.Info("Start applicatiion...")
	return &StdLogger{
		logger: log,
	}
}

func (s *StdLogger) Info(msg string) {
	s.logger.Info(msg)
}

func (s *StdLogger) Debug(msg string) {
	s.logger.Debug(msg)
}

func (s *StdLogger) Error(msg string) {
	s.logger.Error(msg)
}
