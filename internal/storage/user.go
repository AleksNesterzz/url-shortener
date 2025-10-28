package storage

import (
	"fmt"
	"sync"
	"urlshortner/models"
)

type UserRepository interface {
	Register(user models.User) error
	Login(user models.User) error
	GetHashByEmail(user models.User) (string, error)
	//Logout()
}

//adjust logic with models.user

type InMemoryUser struct {
	m     sync.Mutex
	users map[string]string
}

func NewInMemoryUser() *InMemoryUser {
	return &InMemoryUser{
		users: make(map[string]string),
	}
}

func (u *InMemoryUser) Register(user models.User) error {
	u.m.Lock()
	defer u.m.Unlock()
	if _, ok := u.users[user.Email]; ok {
		return fmt.Errorf("user with such email is already exists")
	}
	u.users[user.Email] = user.PasswordHash
	return nil
}

func (u *InMemoryUser) Login(user models.User) error {
	u.m.Lock()
	defer u.m.Unlock()
	_, ok := u.users[user.Email]
	if !ok {
		return fmt.Errorf("such user doesn't exist")
	}
	return nil
}

func (u *InMemoryUser) Logout() {

}

func (a *InMemoryUser) GetHashByEmail(user models.User) (string, error) {
	a.m.Lock()
	defer a.m.Unlock()
	hash, ok := a.users[user.Email]
	if !ok {
		return "", fmt.Errorf("such email didn't exist")
	}
	return hash, nil
}
