package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
)

var Env map[string]string

func SetEnv() {
	var err error
	Env, err = godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetEnv(key, val string) string {
	result := Env[key]
	if result == "" {
		return val
	}
	return result
}
