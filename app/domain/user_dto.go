package domain

import "time"

type UserDto struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
