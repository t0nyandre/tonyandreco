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

	// Configurations, logger and json config
	l := logger.New()
	cfg := config.Load(*appConfig, l)

	l.Info().Msg(fmt.Sprintf("Successfully starting server on port :%v", cfg.Port))
}
