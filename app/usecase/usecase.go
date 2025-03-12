package usecase

import "github.com/fujimisakari/grpc-todo/app/domain/repository"

type Usecase struct {
	logger Logger
	repo   repository.Repository
}

func NewUsecase(logger Logger, repo repository.Repository) *Usecase {
	return &Usecase{
		logger: logger,
		repo:   repo,
	}
}
