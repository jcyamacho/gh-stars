package ports

import (
	"context"

	"github.com/jcyamacho/gh-stars/internal/core/domain"
)

type GitHubClient interface {
	GetUserStarred(context.Context) ([]domain.Repo, error)
}
