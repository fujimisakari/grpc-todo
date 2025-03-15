package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const (
	envLocal      = "local"
	envDevelop    = "develop"
	envProduction = "production"
)

type Environment struct {
	Env      string `envconfig:"ENV" default:"local"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`

	GRPCPort int `envconfig:"GRPC_PORT" default:"5000"`
	HTTPPort int `envconfig:"GRPC_PORT" default:"8000"`

	*Spanner
}

// Spanner stores configuration settings for Google Cloud Spanner.
type Spanner struct {
	ProjectID  string `envconfig:"PROJECT_ID" default:"grpc-todo-spanner-emulator-project"`
	InstanceID string `envconfig:"INSTANCE_ID" default:"grpc-todo-spanner-emulator-instance"`
	DatabaseID string `envconfig:"DATABASE_ID" default:"grpc-todo-spanner-emulator-db"`
}

func (e *Environment) IsLocal() bool {
	return e.Env == envLocal
}

func (e *Environment) IsDevelop() bool {
	return e.Env == envDevelop
}

func (e *Environment) IsProduction() bool {
	return e.Env == envProduction
}

// ReadFromEnv reads configuration from environmental variables
// defined by Environment struct.
func ReadFromEnv() (*Environment, error) {
	env := &Environment{}
	if err := envconfig.Process("", env); err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}

	if err := validate(env.Env); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return env, nil
}

func validate(env string) error {
	checks := []struct {
		bad    bool
		errMsg string
	}{
		{
			env != envLocal && env != envDevelop && env != envProduction,
			fmt.Sprintf("invalid env is specifed: %q", env),
		},

		// Add your own validation here
	}

	for _, check := range checks {
		if check.bad {
			return errors.New(check.errMsg)
		}
	}

	return nil
}
