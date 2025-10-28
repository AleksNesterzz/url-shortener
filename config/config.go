package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type JWTConfig struct {
	SecretKey         string        `env:"JWT_SECRET_KEY,required"`
	AccessTokenExpiry time.Duration `env:"JWT_ACCESS_EXP" default:"1h"`
	Issuer            string        `env:"JWT_ISSUER" default:"url-shortener"`
}

func Load() (*JWTConfig, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	expirationTime, err := time.ParseDuration(os.Getenv("JWT_ACCESS_EXP"))
	if err != nil {
		return nil, err
	}
	return &JWTConfig{
		SecretKey:         os.Getenv("JWT_SECRET_KEY"),
		AccessTokenExpiry: expirationTime,
		Issuer:            os.Getenv("JWT_ISSUER"),
	}, nil
}
