package main

import (
	"flag"

	"forum/config"
	"forum/internal/app"
)

func main() {
	configPath := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	cfg := config.GetConfig(*configPath)

	app.Run(cfg)
}
