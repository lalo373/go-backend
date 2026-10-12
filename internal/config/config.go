package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	ENV   string
	DBURL string
}

func MustLoad() Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT IS REQUIRED")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV IS REQUIRED")
	}
	dburl := os.Getenv("DBURL")
	if dburl == "" {
		panic("DBURL is required")
	}
	return Config{
		Port:  port,
		ENV:   env,
		DBURL: dburl,
	}

}
