package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}
