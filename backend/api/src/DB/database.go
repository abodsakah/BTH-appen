// Package db provides db
package db

import (
	"fmt"
	"log"
	"os"

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
}

// Functions

// SetupDatabase function
// Returns the db session or an error.
func SetupDatabase() (*gorm.DB, error) {
	// get env variables
	var envs dbEnvs
	getDbEnvs(&envs)
	// connect to DB
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Stockholm", envs.User, envs.Pass, envs.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrate database models
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Exam{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&News{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Token{})
	if err != nil {
		return nil, err
	}

	// NOTE: Create admin account so there is something to authenticate with when using the API.
	user := &User{
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
}
