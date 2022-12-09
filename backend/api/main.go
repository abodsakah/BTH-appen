// Package main
package main

import (
	"log"

	"github.com/abodsakah/BTH-appen/backend/api/src/DB"
	"github.com/abodsakah/BTH-appen/backend/api/src/Notifications"
	"github.com/abodsakah/BTH-appen/backend/api/src/Routes"
	"github.com/abodsakah/BTH-appen/backend/api/src/Scraper"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file into process for dev
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("DEV: Could not load .env file")
	}

	gormDB, err := db.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// start scarper
	go scraper.Start(gormDB)

	// start notifications server go routine
	stopRunning := false
	go func() {
		if err := notifications.StartServers(gormDB, &stopRunning, 5); err != nil {
			log.Fatalln("Failed to start expo notifications server, error: ", err)
		}
	}()

	// start API web server main routine
	routes.SetupRoutes(gormDB)
}
