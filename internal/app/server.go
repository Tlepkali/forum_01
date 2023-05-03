package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"forum/config"
	"forum/pkg/logger"
)

func Serve(cfg *config.Config, handler http.Handler) error {
	log := logger.GetLoggerInstance()

	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		log.PrintInfo(fmt.Sprintf("caught signal: %s", s.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		log.PrintInfo("completing background tasks")

		shutdownError <- nil
	}()

	log.PrintInfo(fmt.Sprintf("starting server on %s", srv.Addr))

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	log.PrintInfo("stopped server")

	return nil
}
