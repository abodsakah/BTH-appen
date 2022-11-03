// Package main
package main

import (
	"log"

	"github.com/abodsakah/BTH-appen/backend/api/src/Routes"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file into process for dev
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("DEV: Could not load .env file")
	}

	routes.SetupRoutes()
}
