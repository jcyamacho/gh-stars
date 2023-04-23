package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh stars",
	Short: "gh-starts is a command line tool for managing your GitHub stars",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	ctx := context.Background()
	return rootCmd.ExecuteContext(ctx)
}
