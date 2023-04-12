package app

import (
	"banking-auth/db"
	"banking-auth/env"
)

func Start() {

	config := env.GetConfig()
	_ = db.GetDbClient(config) // use the client to initialize the repositories
}