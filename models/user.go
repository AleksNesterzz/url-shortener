package models

import "time"

type User struct {
	Id           uint
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	IsActive     bool
}
