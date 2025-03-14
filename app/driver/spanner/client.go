package spanner

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"

	"github.com/fujimisakari/grpc-todo/app/driver/config"
)

func NewClient(ctx context.Context, cfg *config.Spanner) (*spanner.Client, error) {
	dbname := fmt.Sprintf("projects/%s/instances/%s/databases/%s", cfg.ProjectID, cfg.InstanceID, cfg.DatabaseID)

	// setup spanner client options

	return spanner.NewClientWithConfig(ctx, dbname,
		spanner.ClientConfig{SessionPoolConfig: spanner.DefaultSessionPoolConfig},
	)
}
