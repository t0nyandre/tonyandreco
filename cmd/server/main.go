package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t0nyandre/tonyandreco/internal/config"
	"github.com/t0nyandre/tonyandreco/internal/logger"
)

var appConfig = flag.String("config", "./config/dev.json", "path to config file")

func main() {
	flag.Parse()

	// Configurations, logger and json config
	l := logger.New()
	cfg := config.Load(*appConfig, l)

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	l.Info().Msgf("Server listening on port %v", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%v", cfg.Hostname, cfg.Port), router); err != nil {
		l.Fatal().Err(err).Msg("error starting server")
	}
}
