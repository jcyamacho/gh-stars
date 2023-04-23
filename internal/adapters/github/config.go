package github

import "github.com/cli/go-gh/v2/pkg/api"

type Config struct {
	Client *api.RESTClient
}

func configDefault(config ...Config) (Config, error) {
	var cfg Config

	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.Client == nil {
		client, err := api.DefaultRESTClient()
		if err != nil {
			return cfg, err
		}
		cfg.Client = client
	}

	return cfg, nil
}
