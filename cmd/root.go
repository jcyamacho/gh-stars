package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	errCh := make(chan error, 1)

	go func() {
		errCh <- rootCmd.ExecuteContext(ctx)
	}()

	select {
	case <-c:
		cancel()
	case err := <-errCh:
		return err
	}

	return <-errCh
}
