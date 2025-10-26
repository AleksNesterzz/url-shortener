package storage

import (
	"fmt"
	"sync"
)

type UserRepository interface {
	Register(email, password string) error
	Login(email, pass string) error
	//Logout()
}

type InMemoryUser struct {
	m     sync.Mutex
	users map[string]string
}

func NewInMemoryUser() *InMemoryUser {
	return &InMemoryUser{
		users: make(map[string]string),
	}
}

func (u *InMemoryUser) Register(email, password string) error {
	u.m.Lock()
	defer u.m.Unlock()
	if _, ok := u.users[email]; ok {
		return fmt.Errorf("user with such email is already exists")
	}
	u.users[email] = password
	return nil
}

func (u *InMemoryUser) Login(email, pass string) error {
	u.m.Lock()
	defer u.m.Unlock()
	hash, ok := u.users[email]
	if !ok {
		return fmt.Errorf("such user doesn't exist")
	}
	if hash != pass {
		return fmt.Errorf("password is incorrect")
	}
	return nil
}

func (u *InMemoryUser) Logout() {

}
