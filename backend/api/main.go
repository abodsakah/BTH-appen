// Package main
package main

import (
	"log"

	//routes "github.com/abodsakah/BTH-appen/backend/api/src/Routes"
	scraper "github.com/abodsakah/BTH-appen/backend/api/src/Scraper"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file into process for dev
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("DEV: Could not load .env file")
	}

	scraper.GetNews()
	//routes.SetupRoutes()
}
