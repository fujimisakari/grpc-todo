package domain

import "time"

type TodoDto struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Priority    Priority
	Completed   bool
	DueTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
