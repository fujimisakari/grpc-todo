package repository

import (
	"github.com/fujimisakari/grpc-todo/internal/domain/repository"
)

// Repository is an interface for repository.
type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Todo() repository.Todo {
	return &TodoRepository{}
}
