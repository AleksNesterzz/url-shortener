package storage

type UserRepository interface {
	Register()
	Login()
	Logout()
}
