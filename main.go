package main

import (
	app "go-bank-api/app"
	logger "go-bank-api/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
