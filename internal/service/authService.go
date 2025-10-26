package service

import (
	"urlshortner/internal/logger"
	"urlshortner/internal/storage"
)

type Auth struct {
	logger logger.Logger
	user   storage.UserRepository
}

func NewAuth(u storage.UserRepository, logger logger.Logger) *Auth {
	return &Auth{
		user: u, 
		logger: logger
	}
}

func (a *Auth) Register(email string, password string) error {
	a.logger.Info("registering user: " + email)
	err := a.user.Register(email, password)
	if err!=nil {
		return err
	}
	return nil
}

func (a *Auth) Login(email string, pass string) error  {
	a.logger.Info("logging in user: " + email)
	err := a.user.Login(email, pass)
	if err!=nil {
		return err
	}
	return nil
}

// func (a *Auth) Logout() {
// 	//a.user.Logout()
// }
