package github

import "github.com/cli/go-gh/v2/pkg/api"

type Client struct {
	rc *api.RESTClient
}

func NewClient(config ...Config) (*Client, error) {
	cfg, err := configDefault(config...)
	if err != nil {
		return nil, err
	}

	return &Client{
		rc: cfg.Client,
	}, nil
}
