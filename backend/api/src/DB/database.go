// Package db provides db
package db

import (
	"fmt"
	"log"
	"os"

	models "github.com/abodsakah/BTH-appen/backend/api/src/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// constants
const (
	passMaxLength int = 50
	userMaxLength int = 20
)

// structs
type dbEnvs struct {
	Name string
	User string
	Pass string
	Host string
}

// Functions

// SetupDatabase function
// Returns the db session or an error.
func SetupDatabase() (*gorm.DB, error) {
	// get env variables
	var envs dbEnvs
	getDbEnvs(&envs)
	// connect to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Stockholm", envs.Host, envs.User, envs.Pass, envs.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrate database models
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Exam{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.News{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Token{})
	if err != nil {
		return nil, err
	}

	// NOTE: Create admin account so there is something to authenticate with when using the API.
	user := &models.User{
		Name:     "Admin Adminsson",
		Username: "admin",
		Password: "pass",
		Role:     "admin",
	}
	err = CreateUser(db, user)
	if err != nil {
		log.Println(err)
	}
	// return db object
	return db, nil
}

// Takes a dbEnvs struct pointer and sets all the env vars.
func getDbEnvs(dbEnvs *dbEnvs) {
	var ok bool

	dbEnvs.Name, ok = os.LookupEnv("POSTGRES_NAME")
	if !ok {
		log.Fatalln("POSTGRES_NAME: env variable not found")
	}

	dbEnvs.User, ok = os.LookupEnv("POSTGRES_USER")
	if !ok {
		log.Fatalln("POSTGRES_USER: env variable not found")
	}

	dbEnvs.Pass, ok = os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		log.Fatalln("POSTGRES_PASSWORD: env variable not found")
	}
	dbEnvs.Host, ok = os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatalln("DB_HOST: env variable not found")
	}
}
