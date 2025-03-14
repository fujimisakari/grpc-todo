package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	"cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
)

func main() {
	var code int
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		code = 1
	}
	os.Exit(code)
}

func run() error {
	// For safety
	if emu := os.Getenv("SPANNER_EMULATOR_HOST"); emu == "" {
		return errors.New("start spanner emulator and set $SPANNER_EMULATOR_HOST")
	}

	projectID := os.Getenv("SPANNER_PROJECT")
	if projectID == "" {
		return errors.New("$SPANNER_PROJECT must be set")
	}

	instanceID := os.Getenv("SPANNER_INSTANCE")
	if instanceID == "" {
		return errors.New("$SPANNER_INSTANCE must be set")
	}

	ctx := context.Background()

	instanceAdmin, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = instanceAdmin.Close()
	}()

	op, err := instanceAdmin.CreateInstance(ctx, &instancepb.CreateInstanceRequest{
		Parent:     fmt.Sprintf("projects/%s", projectID),
		InstanceId: instanceID,
		Instance: &instancepb.Instance{
			Config:      fmt.Sprintf("projects/%s/instanceConfigs/emulator-config", projectID),
			DisplayName: instanceID,
			NodeCount:   1,
		},
	})
	if err != nil {
		return fmt.Errorf("could not create instance projects/%s/instances/%s: %w", projectID, instanceID, err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for instance creation to finish failed: %w", err)
	}

	return nil
}
