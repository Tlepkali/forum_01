package app

import (
	"fmt"

	"forum/config"
	"forum/internal/handler"
	"forum/internal/render"
	"forum/internal/repository"
	"forum/internal/service"
	"forum/pkg/client/sqlite"
	"forum/pkg/logger"
)

func Run(cfg *config.Config) {
	log := logger.GetLoggerInstance()

	db, err := sqlite.OpenDB(cfg.DB.DSN)
	if err != nil {
		log.PrintFatal(err)
	}
	log.PrintInfo("Connected to DB")

	repo := repository.NewRepository(db)
	log.PrintInfo("Connected to Repository")

	service := service.NewService(repo)
	log.PrintInfo("Connected to Service")

	templateCache, err := render.NewTemplateCache(cfg.TemplatesDir)
	if err != nil {
		log.PrintFatal(err)
	}

	handler := handler.NewHandler(service, templateCache)
	log.PrintInfo("Connected to Handler")

	fmt.Println("Server is running on port", cfg.Addr)

	if err := Serve(cfg, handler.Routes(cfg)); err != nil {
		log.PrintFatal(err)
	}
}
