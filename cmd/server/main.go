package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/t0nyandre/tonyandreco/internal/config"
)

var appConfig = flag.String("config", "./config/dev.json", "path to config file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*appConfig)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Successfully loaded config for %s\n", cfg.Defaults.AppName)
}
