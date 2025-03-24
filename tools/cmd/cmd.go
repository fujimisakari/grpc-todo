//go:build cmd
// +build cmd

package cmd

import (
	_ "github.com/cloudspannerecosystem/spanner-cli"
	_ "github.com/cloudspannerecosystem/wrench"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/incu6us/goimports-reviser/v3"
	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
	_ "go.mercari.io/yo"
	_ "go.uber.org/mock/mockgen"
)
