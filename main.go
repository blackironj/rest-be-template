package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/blackironj/rest-be-template/env"
	"github.com/blackironj/rest-be-template/repository"
	"github.com/blackironj/rest-be-template/server"
	"github.com/blackironj/rest-be-template/util/logger"
)

func main() {
	env.Init()
	logger.Init()
	repository.Init()

	app := server.Init()

	go func() {
		if err := app.Listen(":" + env.SrvPort); err != nil {
			log.Err(err).Send()
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // This blocks the main thread until an interrupt is received
	log.Info().Msg("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Info().Msg("Running cleanup tasks...")
	repository.CloseMongoDB()
	log.Info().Msg("Server was successful shutdown.")
}
