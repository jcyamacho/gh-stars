package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jcyamacho/gh-stars/internal/core/ports"
)

type UserStarredDownloader struct {
	client ports.GitHubClient
}

func NewUserStarredDownloader(client ports.GitHubClient) *UserStarredDownloader {
	return &UserStarredDownloader{client: client}
}

func (u *UserStarredDownloader) Download(ctx context.Context, w io.Writer) error {
	log.Println("downloading user starred...")

	repos, err := u.client.GetUserStarred(ctx)
	if err != nil {
		return fmt.Errorf("error getting user starred: %w", err)
	}

	log.Println("encoding...")

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	return enc.Encode(repos)
}
