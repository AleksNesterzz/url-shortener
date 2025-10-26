package models

import "time"

type User struct {
	Id           uint64
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	IsActive     bool
}
