package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jcyamacho/gh-stars/internal/adapters/github"
	"github.com/jcyamacho/gh-stars/internal/core/usecases"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download [file]",
	Short: "Download your GitHub stars",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName := "starred-repos.json"
		if len(args) > 0 {
			fileName = args[0]
		}
		if !strings.HasSuffix(fileName, ".json") {
			fileName += ".json"
		}

		f, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}

		defer f.Close()

		client, err := github.NewClient()
		if err != nil {
			return fmt.Errorf("error creating github client: %w", err)
		}

		downloader := usecases.NewUserStarredDownloader(client)

		if err := downloader.Download(cmd.Context(), f); err != nil {
			return fmt.Errorf("error downloading starred repos: %w", err)
		}

		log.Printf("downloaded starred repos to: %s\n", fileName)

		return nil
	},
}
