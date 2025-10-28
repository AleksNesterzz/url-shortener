package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
	"urlshortner/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenGenerator interface {
	GenerateAccessToken(*models.User) (string, error)
	ValidateToken(string) (*models.TokenClaims, error)
	ParseTokenClaims(string) (*models.TokenClaims, error)
	//GenerateRefreshToken(*models.User) (string, error)
}

type JWTGenerator struct {
	secretKey          []byte
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
	issuer             string
}

type JWTClaims struct {
	UserID uint             `json:"user_id"`
	Email  string           `json:"email"`
	Type   models.TokenType `json:"type"`
	jwt.RegisteredClaims
}

func NewJWTGenerator(secretKey string, accessTokenExpiry, refreshTokenExpiry time.Duration, issuer string) (*JWTGenerator, error) {
	if secretKey == "" {
		return nil, fmt.Errorf("secret key cannot be empty")
	}
	return &JWTGenerator{
		secretKey:          []byte(secretKey),
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
		issuer:             issuer,
	}, nil
}

// hide in .env
var secretKey = []byte("mysecret")

func (g *JWTGenerator) GenerateAccessToken(user *models.User) (string, error) {
	if user == nil {
		return "", fmt.Errorf("user cannot be nil")
	}
	expirationTime := time.Now().Add(g.accessTokenExpiry)
	claims := &JWTClaims{
		UserID: user.Id,
		Email:  user.Email,
		Type:   models.TokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    g.issuer,
			ID:        generateTokenID(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(g.secretKey)
}

func (g *JWTGenerator) VerifyToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

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

func (g *JWTGenerator) GenerateRefreshToken(user models.User) (string, error) {
	return "", nil
}

func (g *JWTGenerator) ParseTokenClaims(tokenString string) (*models.TokenClaims, error) {
	claims, err := g.parseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return g.toModelsClaims(claims), nil
}

func (g *JWTGenerator) parseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return g.secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token parsing fialed: %w", err)
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func (g *JWTGenerator) toModelsClaims(jwtClaims *JWTClaims) *models.TokenClaims {
	return &models.TokenClaims{
		UserID:    jwtClaims.UserID,
		Email:     jwtClaims.Email,
		Type:      jwtClaims.Type,
		IssuedAt:  jwtClaims.IssuedAt.Time,
		ExpiresAt: jwtClaims.ExpiresAt.Time,
	}
}

func generateTokenID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return uuid.New().String()
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
