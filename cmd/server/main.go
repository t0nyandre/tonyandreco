package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/t0nyandre/tonyandreco/internal/config"
	"github.com/t0nyandre/tonyandreco/internal/logger"
	"github.com/t0nyandre/tonyandreco/internal/routes"
)

var appConfig = flag.String("config", "./config/dev.json", "path to config file")

func main() {
	flag.Parse()

	// Configurations, logger and json config
	l := logger.New()
	if err := config.Load(*appConfig); err != nil {
		l.Fatal().Err(err).Msg("error loading config")
	}
	cfg := config.AppConfig

	r := routes.NewRouter(l)

	l.Info().Msgf("Server listening on port %v", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%v", cfg.Hostname, cfg.Port), r); err != nil {
		l.Fatal().Err(err).Msg("error starting server")
	}
}
