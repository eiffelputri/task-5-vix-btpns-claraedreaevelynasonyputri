package app

import (
	"time"
)

type User struct {
	ID        uint
	Username  string
	Email     string
	Password  byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
