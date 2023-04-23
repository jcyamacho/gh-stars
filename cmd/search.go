package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jcyamacho/gh-stars/internal/adapters/github"
	"github.com/jcyamacho/gh-stars/internal/core/services"
	"github.com/jcyamacho/gh-stars/internal/core/usecases"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search [text]",
	Short: "Search your GitHub stars",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := strings.Join(args, " ")

		client, err := github.NewClient()
		if err != nil {
			return fmt.Errorf("error creating github client: %w", err)
		}

		searcher := usecases.NewUserStarredSearcher(client, services.TextSearchFactory{})
		repos, err := searcher.Search(cmd.Context(), text)
		if err != nil {
			return fmt.Errorf("error searching starred repos: %w", err)
		}

		if len(repos) == 0 {
			fmt.Println("no results found")
			return nil
		}

		log.Printf("found %d repos:", len(repos))

		for i, repo := range repos {
			u := repo.Homepage
			if u == "" {
				u = repo.URL
			}

			fmt.Printf("%2d. [%s]: %s\n", i+1, repo.FullName, u)
		}

		return nil
	},
}
