package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

func parseEnvConfig(config Config) Config {
	ctx := context.Background()

	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
