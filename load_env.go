package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	GoEnv string
	Port  string
}

func LoadEnv() Environment {
	goEnv := os.Getenv("GO_ENV")

	if goEnv != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")

	env := Environment{
		GoEnv: goEnv,
		Port:  port,
	}

	return env
}
