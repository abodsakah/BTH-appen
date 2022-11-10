// Package db provides db
package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// constants
const passMaxLength int = 50
const userMaxLength int = 20

// variables
var db *gorm.DB

var (
	err        error
	dbName     string
	dbUser     string
	dbPassword string
)

// Functions

// SetupDatabase function
// Returns the db session or an error.
func SetupDatabase() (*gorm.DB, error) {
	// get env variables
	getEnvs()
	// connect to DB
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Stockholm", dbUser, dbPassword, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrate database models
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Exam{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&News{})
	if err != nil {
		log.Println(err)
	}

	// create admin account
	user := &User{Username: "admin", Password: "pass"}
	err = CreateUser(user)
	if err != nil {
		log.Println(err)
	}

	// get first user
	var userOne User
	db.First(&userOne)

	// create some exams
	err = CreateExam(&Exam{
		CourseCode: "DV1337",
		StartDate:  time.Now().AddDate(0, 2, 0),
		Users:      []User{userOne},
	})
	if err != nil {
		log.Println(err)
	}
	err = CreateExam(&Exam{
		CourseCode: "PA6969",
		StartDate:  time.Now().AddDate(0, 2, 5),
		Users:      []User{userOne},
	})
	if err != nil {
		log.Println(err)
	}

	// list exams
	exams, err := ListExams()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("\nExams: %#v\n", exams)

	return db, nil
}

// gets env variables and puts them in appropriate variables.
func getEnvs() {
	var ok bool

	dbName, ok = os.LookupEnv("POSTGRES_NAME")
	if !ok {
		log.Fatalln("POSTGRES_NAME: env variable not found")
	}

	dbUser, ok = os.LookupEnv("POSTGRES_USER")
	if !ok {
		log.Fatalln("POSTGRES_USER: env variable not found")
	}

	dbPassword, ok = os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		log.Fatalln("POSTGRES_PASSWORD: env variable not found")
	}
}
