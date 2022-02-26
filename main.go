package main

import (
	"github.com/rtpa25/banking/app"
	"github.com/rtpa25/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
