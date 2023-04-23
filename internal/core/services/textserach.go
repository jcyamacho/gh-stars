package services

import (
	"fmt"
	"sync"

	"github.com/blevesearch/bleve/v2"
	"github.com/jcyamacho/gh-stars/internal/core/domain"
	"github.com/jcyamacho/gh-stars/internal/core/ports"
)

type TextSearch struct {
	index bleve.Index
	repos map[string]domain.Repo
	mx    sync.Mutex
}

func NewTextSearch() (*TextSearch, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return nil, fmt.Errorf("error creating index: %w", err)
	}

	return &TextSearch{
		index: index,
		repos: make(map[string]domain.Repo),
	}, nil
}

func (t *TextSearch) Index(r domain.Repo) error {
	t.mx.Lock()
	defer t.mx.Unlock()

	t.repos[r.FullName] = r

	return t.index.Index(r.FullName, map[string]any{
		"name":        r.Name,
		"description": r.Description,
		"topics":      r.Topics,
	})
}

func (t *TextSearch) Search(text string) ([]domain.Repo, error) {
	query := bleve.NewMatchQuery(text)
	search := bleve.NewSearchRequest(query)

	result, err := t.index.Search(search)
	if err != nil {
		return nil, fmt.Errorf("error searching: %w", err)
	}

	results := make([]domain.Repo, len(result.Hits))

	for i, hit := range result.Hits {
		results[i] = t.repos[hit.ID]
	}

	return results, nil
}

type TextSearchFactory struct{}

func (TextSearchFactory) NewTextSearch() (ports.TextSearch, error) {
	return NewTextSearch()
}
