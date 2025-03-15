package domain

import "time"

type Todo struct {
	ID          string
	Title       string
	Description string
	Priority    Priority
	Completed   bool
	DueTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
