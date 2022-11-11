package main

import (
	"demi-anchor/internal/config"
	"demi-anchor/internal/handler"
	"demi-anchor/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.InitLogger(cfg.LogLevel, cfg.IsJsonLog)

	log.Info().Msg("Starting")

	h := handler.New()

	go func() {
		err = http.ListenAndServe(":"+cfg.Port, h.InitRouter())
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	log.Info().Msg("Shutdown...")

	// Clean/close/kill ...

	log.Info().Msg("Goodbye!")
}
