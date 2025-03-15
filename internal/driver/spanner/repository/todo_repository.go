package repository

import (
	"context"
	"strings"

	"cloud.google.com/go/spanner"

	"github.com/fujimisakari/grpc-todo/internal/domain"
	domain_repo "github.com/fujimisakari/grpc-todo/internal/domain/repository"
)

// FindByID gets a todo by id.
func (t *TodoRepository) FindByID(ctx context.Context, ro domain_repo.YORODB, id string) (*domain.Todo, error) {
	return t.Find(ctx, ro, id)
}

// ListTodo lists todos.
func (t *TodoRepository) ListTodo(ctx context.Context, ro domain_repo.YORODB) ([]*domain.Todo, error) {
	sqlstr := `SELECT ` +
		strings.Join(t.columns(), ",") +
		` FROM ` + TodoTableName +
		` Limit 20`

	stmt := spanner.NewStatement(sqlstr)
	return t.FindByStatement(ctx, ro, stmt)
}
