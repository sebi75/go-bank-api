package main

import (
	app "go-bank-api/app"
	"go-bank-api/env"
	logger "go-bank-api/logger"
)

func main() {
	env.SanityCheck()
	logger.Info("Starting the application...")
	app.Start()
}
