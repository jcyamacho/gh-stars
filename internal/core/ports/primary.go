package ports

import "github.com/jcyamacho/gh-stars/internal/core/domain"

type TextSearch interface {
	Index(repo domain.Repo) error
	Search(text string) ([]domain.Repo, error)
}

type TextSearchFactory interface {
	NewTextSearch() (TextSearch, error)
}
