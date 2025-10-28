package token

import (
	"fmt"
	"time"
	"urlshortner/models"

	"github.com/golang-jwt/jwt/v5"
)

type TokenRepository interface {
	GenerateAccessToken(models.User) (string, error)
	ValidateToken(string) (*Claims, error)
	GenerateRefreshToken(models.User) (string, error)
}

type TokenGenerator struct {
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// hide in .env
var secretKey = []byte("mysecret")

func (t *TokenGenerator) GenerateAccessToken(user models.User) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.Id,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "urlshort",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func (t *TokenGenerator) GenerateRefreshToken(user models.User) (string, error) {
	return "", nil
}
