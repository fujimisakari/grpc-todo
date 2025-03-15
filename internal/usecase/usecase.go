package usecase

import (
	"cloud.google.com/go/spanner"

	"github.com/fujimisakari/grpc-todo/internal/domain/repository"
)

type Usecase struct {
	logger        Logger
	spannerClient *spanner.Client
	repo          repository.Repository
}

func NewUsecase(logger Logger, spannerClient *spanner.Client, repo repository.Repository) *Usecase {
	return &Usecase{
		logger:        logger,
		spannerClient: spannerClient,
		repo:          repo,
	}
}
