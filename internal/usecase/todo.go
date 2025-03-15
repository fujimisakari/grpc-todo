package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/fujimisakari/grpc-todo/internal/domain"
)

func (u *Usecase) GetTodo(ctx context.Context, todoID string) (*domain.Todo, error) {
	return u.repo.Todo().GetTodoByID(ctx, todoID)
}

func (u *Usecase) ListTodo(ctx context.Context) ([]*domain.Todo, error) {
	return u.repo.Todo().ListTodo(ctx)
}

func (u *Usecase) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	todo.ID = uuid.New().String()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	return u.repo.Todo().CreateTodo(ctx, todo)
}

func (u *Usecase) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	return u.repo.Todo().UpdateTodo(ctx, todo)
}

func (u *Usecase) UpdateTodoStatus(ctx context.Context, todoID string, completed bool) error {
	todo, err := u.repo.Todo().GetTodoByID(ctx, todoID)
	if err != nil {
		return err
	}
	todo.Completed = completed
	return u.repo.Todo().UpdateTodo(ctx, todo)
}

func (u *Usecase) DeleteTodo(ctx context.Context, todoID string) error {
	return u.repo.Todo().DeleteTodo(ctx, todoID)
}
