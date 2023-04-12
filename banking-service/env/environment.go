package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("Loading environment variables...")
	godotenv.Load()
	log.Println("Environment variables loaded")
}

type Config struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

func GetConfig() *Config {
	return &Config{
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}
}

func SanityCheck() {
	log.Println("Sanity checking environment variables...")
	if GetConfig().DbHost == "" {
		log.Fatal("DB_HOST is not set")
	}
	if GetConfig().DbPort == "" {
		log.Fatal("DB_PORT is not set")
	}
	if GetConfig().DbUser == "" {
		log.Fatal("DB_USER is not set")
	}
	if GetConfig().DbPass == "" {
		log.Fatal("DB_PASS is not set")
	}
	if GetConfig().DbName == "" {
		log.Fatal("DB_NAME is not set")
	}
	log.Println("Sanity check passed")
}