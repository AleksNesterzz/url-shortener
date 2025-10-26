package service

import "urlshortner/internal/storage"

type Auth struct {
	user storage.UserRepository
}

func NewAuth(u storage.UserRepository) *Auth {
	return &Auth{user: u}
}

func (a *Auth) Register(email string, password string) {
	a.user.Register(email, password)
}

func (a *Auth) Login(email string, pass string) {
	a.user.Login(email, pass)
}

// func (a *Auth) Logout() {
// 	//a.user.Logout()
// }
