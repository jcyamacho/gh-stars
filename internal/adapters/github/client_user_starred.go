package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jcyamacho/gh-stars/internal/core/domain"
)

const PageSize = 100

func (c *Client) GetUserStarred(ctx context.Context) ([]domain.Repo, error) {
	result := make([]domain.Repo, 0, PageSize)

	size := PageSize
	page := 1

	for size == PageSize {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			data := make([]domain.Repo, 0, PageSize)
			path := fmt.Sprintf("user/starred?per_page=%d&page=%d", PageSize, page)

			err := c.rc.DoWithContext(ctx, http.MethodGet, path, nil, &data)
			if err != nil {
				return nil, fmt.Errorf("error getting user starred: %w", err)
			}

			result = append(result, data...)

			size = len(data)
			page++
		}
	}

	return result, nil
}
