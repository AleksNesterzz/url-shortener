package config

import "time"

type JWTConfig struct {
	SecretKey         string        `env:"JWT_SECRET_KEY,required"`
	AccessTokenExpiry time.Duration `env:"JWT_ACCESS_EXP" default:"1h"`
	Issuer            string        `env:"JWT_ISSUER" default:"url-shortener"`
}
