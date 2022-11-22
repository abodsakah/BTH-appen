// Package main
package main

import (
	"log"

	//routes "github.com/abodsakah/BTH-appen/backend/api/src/Routes"
	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	notifications "github.com/abodsakah/BTH-appen/backend/api/src/Notifications"
	routes "github.com/abodsakah/BTH-appen/backend/api/src/Routes"
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
	gormDB, err := db.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// start notifications server go routine
	go func() {
		if err := notifications.StartServer(gormDB); err != nil {
			log.Fatalln("Failed to start expo notifications server, error: ", err)
		}
	}()

	// start API web server main routine
	routes.SetupRoutes(gormDB)
}
