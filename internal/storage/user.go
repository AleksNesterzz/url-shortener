package storage

type UserRepository interface {
	Register()
	Login()
	Logout()
}

type InMemoryUser struct {
	users map[string]string
}

func NewInMemoryUser() *InMemoryUser {
	return &InMemoryUser{
		users: make(map[string]string),
	}
}

func (u *InMemoryUser) Register() {

}

func (u *InMemoryUser) Login() {

}

func (u *InMemoryUser) Logout() {

}
