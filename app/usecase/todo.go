package usecase

import (
	"context"
	"time"

	"github.com/fujimisakari/grpc-todo/app/domain"
)

func (u *Usecase) GetTodo(ctx context.Context, todoID string) (*domain.Todo, error) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	return &domain.Todo{
		ID:          "ID",
		UserID:      "userid",
		Title:       "title",
		Description: "description",
		Priority:    domain.PRIORITY_HIGH,
		Completed:   false,
		DueTime:     now,
		CreatedAt:   yesterday,
	}, nil
}
