package service

import "urlshortner/internal/storage"

type Auth struct {
	user storage.UserRepository
}

func NewAuth(u storage.UserRepository) *Auth {
	return &Auth{user: u}
}

func (a *Auth) Register() {

}

func (a *Auth) Login() {

}

func (a *Auth) Logout() {

}
