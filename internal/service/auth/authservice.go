package service

import (
	"urlshortner/internal/logger"
	"urlshortner/internal/storage"
	"urlshortner/internal/token"
	"urlshortner/models"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	logger         logger.Logger
	user           storage.UserRepository
	tokenGenerator token.TokenGenerator
}

func NewAuth(u storage.UserRepository, logger logger.Logger, token token.TokenGenerator) *Auth {
	return &Auth{
		user:           u,
		logger:         logger,
		tokenGenerator: token,
	}
}

func (a *Auth) Register(email string, password string) error {
	a.logger.Info("registering user: " + email)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		a.logger.Error("gen password error(registering)" + err.Error())
		return err
	}
	user := models.User{
		Email:        email,
		PasswordHash: string(hash),
	}
	err = a.user.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) Login(email string, pass string) (string, error) {
	a.logger.Info("logging in user: " + email)
	user := models.User{
		Email: email,
	}
	hash, err := a.user.GetHashByEmail(user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		a.logger.Error("password not correct:" + err.Error())
	}

	token, err := a.tokenGenerator.GenerateAccessToken(&user)
	if err != nil {
		a.logger.Error("jwt generating error:" + err.Error())
		return "", err
	}
	err = a.user.Login(user)
	if err != nil {
		a.logger.Error("login error occured:" + err.Error())
		return "", err
	}
	return token, nil
}

// func (a *Auth) Logout() {
// 	//a.user.Logout()
// }
