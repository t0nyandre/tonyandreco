package main

import (
	"flag"
	"fmt"

	"github.com/t0nyandre/tonyandreco/internal/config"
	"github.com/t0nyandre/tonyandreco/internal/logger"
)

var appConfig = flag.String("config", "./config/dev.json", "path to config file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*appConfig)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %v", err))
	}

	logs, err := logger.New(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %v", err))
	}

	logs.Info().Msg("starting server")
}
