package main

import (
	"demi-anchor/internal/config"
	"demi-anchor/internal/handler"
	"demi-anchor/internal/repository"
	"demi-anchor/internal/service"
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

	logger.Init(cfg.LogLevel, cfg.IsJsonLog)

	log.Info().Msg("Starting")

	r, err := repository.New(cfg.Repository)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	s := service.New(r)
	h := handler.New(s)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":"+cfg.Port, h.InitRouter()); err != nil {
			log.Err(err).Send()
			sigChan <- syscall.SIGTERM
		}
	}()

	<-sigChan

	log.Info().Msg("Shutdown...")

	r.Close()

	log.Info().Msg("Goodbye!")
}
