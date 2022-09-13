package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackironj/rest-be-template/env"
	"github.com/blackironj/rest-be-template/server"
	"github.com/blackironj/rest-be-template/util/logger"
)

func main() {
	logger.Init()

	app := server.Init()

	go func() {
		if err := app.Listen(":" + env.SrvPort); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // This blocks the main thread until an interrupt is received
	logger.Debug("Gracefully shutting down...")
	_ = app.Shutdown()

	logger.Debug("Running cleanup tasks...")
	//TODO: close db...

	logger.Debug("Server was successful shutdown.")
}
