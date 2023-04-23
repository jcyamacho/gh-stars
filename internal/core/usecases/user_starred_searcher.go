package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/jcyamacho/gh-stars/internal/core/domain"
	"github.com/jcyamacho/gh-stars/internal/core/ports"
)

type UserStarredSearcher struct {
	client    ports.GitHubClient
	tsFactory ports.TextSearchFactory
}

func NewUserStarredSearcher(client ports.GitHubClient, tsFactory ports.TextSearchFactory) *UserStarredSearcher {
	return &UserStarredSearcher{client: client, tsFactory: tsFactory}
}

func (u *UserStarredSearcher) Search(ctx context.Context, text string) ([]domain.Repo, error) {
	log.Println("downloading user starred...")

	repos, err := u.client.GetUserStarred(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting user starred: %w", err)
	}

	ts, err := u.tsFactory.NewTextSearch()
	if err != nil {
		return nil, fmt.Errorf("error creating text searcher: %w", err)
	}

	log.Println("indexing repos...")

	for _, repo := range repos {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if err := ts.Index(repo); err != nil {
				return nil, fmt.Errorf("error indexing repo: %w", err)
			}
		}
	}

	log.Println("searching...")

	return ts.Search(text)
}
