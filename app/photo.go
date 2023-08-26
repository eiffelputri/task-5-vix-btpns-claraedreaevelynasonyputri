package app

import (
	"time"
)

type Photo struct {
	ID        uint
	Title     string
	Caption   string
	PhotoURL  string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
