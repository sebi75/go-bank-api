package db

import (
	"go-bank-api/env"
	"go-bank-api/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetDbClient(config *env.Config) *sqlx.DB {
	connectionString := config.DbUser + ":" + config.DbPass + "@tcp(" + config.DbHost + ":" + config.DbPort + ")/" + config.DbName
	client, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	logger.Info("Database connection successful")

	return client
}