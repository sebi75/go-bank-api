package main

import (
	"banking-auth/app"
	"banking-auth/env"
	"banking-auth/logger"
)

func main() {
	env.SanityCheck()
	logger.Info("Starting the application...")
	// here will come the code from the app to start the server
	app.Start()
}