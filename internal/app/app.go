package app

import (
	"github.com/spf13/cobra"

	"github.com/far4599/swagger-openapiv2-merge/internal/commands"
)

type app struct {
	rootCmd *cobra.Command
}

func NewApp() *app {
	rootCmd := &cobra.Command{
		Use:   "swagger-tools",
		Short: "Set of tools to work with swagger files.",
	}

	rootCmd.AddCommand(commands.NewMergeCommand())
	rootCmd.AddCommand(commands.NewServeCommand())

	return &app{
		rootCmd: rootCmd,
	}
}

func (a app) Run() error {
	return a.rootCmd.Execute()
}
