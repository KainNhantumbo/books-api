package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env value")
	}
	// Return the value of the variable
	return os.Getenv(key)
}
