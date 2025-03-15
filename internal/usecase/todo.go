package usecase

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"

	"github.com/fujimisakari/grpc-todo/internal/domain"
)

func (u *Usecase) GetTodo(ctx context.Context, todoID string) (*domain.Todo, error) {
	return u.repo.Todo().FindByID(ctx, u.spannerClient.Single(), todoID)
}

func (u *Usecase) ListTodo(ctx context.Context) ([]*domain.Todo, error) {
	return u.repo.Todo().ListTodo(ctx, u.spannerClient.Single())
}

func (u *Usecase) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	_, err := u.spannerClient.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		todo.ID = uuid.New().String()
		return tx.BufferWrite([]*spanner.Mutation{u.repo.Todo().Insert(todo)})
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	_, err := u.spannerClient.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		t, err := u.repo.Todo().FindByID(ctx, u.spannerClient.Single(), todo.ID)
		if err != nil {
			return err
		}
		t.Title = todo.Title
		t.Description = todo.Description
		t.Priority = todo.Priority
		return tx.BufferWrite([]*spanner.Mutation{u.repo.Todo().Update(t)})
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) UpdateTodoStatus(ctx context.Context, todoID string, completed bool) error {
	_, err := u.spannerClient.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		t, err := u.repo.Todo().FindByID(ctx, u.spannerClient.Single(), todoID)
		if err != nil {
			return err
		}
		t.Completed = completed
		return tx.BufferWrite([]*spanner.Mutation{u.repo.Todo().Update(t)})
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) DeleteTodo(ctx context.Context, todoID string) error {
	_, err := u.spannerClient.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		t, err := u.repo.Todo().FindByID(ctx, u.spannerClient.Single(), todoID)
		if err != nil {
			return err
		}
		return tx.BufferWrite([]*spanner.Mutation{u.repo.Todo().Delete(t)})
	})
	if err != nil {
		return err
	}
	return nil
}
