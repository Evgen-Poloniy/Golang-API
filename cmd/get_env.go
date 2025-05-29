package main

import (
	"os"

	"github.com/joho/godotenv"
)

func getEnv(envPath string) error {
	_, err := os.Stat("/.dockerenv")
	if err != nil {
		err = godotenv.Load(envPath)
	}

	return err
}
