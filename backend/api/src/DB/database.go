// Package db provides db
package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
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

// SetupDatabase function
func SetupDatabase() {
	// get env variables
	getEnvs()
	// connect to DB
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Stockholm", dbUser, dbPassword, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// migrate database models
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Exam{})

	// create admin account
	user := &User{Username: "admin", Password: "pass"}
	err = CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	// create some commands
	exam := &Exam{
		CourseCode: "DV1337",
		StartDate:  time.Now(),
		Users:      []User{*user},
	}
	err := CreateExam(exam)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Command: %#v\n", exam)

	commands, err := ListCommands()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\nCommands: %#v\n", commands)
}

func checkInputLength(username string, password string) (err error) {
	if len(username) > userMaxLength || len(password) > passMaxLength {
		return errors.New("Error: username or password exceeds max length")
	}
	return nil
}

// AuthUser function
func AuthUser(username string, password string) (userID string, err error) {
	// check username and password length
	err = checkInputLength(username, password)
	if err != nil {
		return "", err
	}

	// get user from database
	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	// compare password and hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err // return if password doesn't match
	}

	// return user ID
	userID = strconv.Itoa(int(user.ID))
	return userID, nil
}

// CreateExam function
func CreateExam(exam *Exam) error {
	// set creation date
	exam.CreatedAt = time.Now()

	// create exam in database
	result := db.Create(&exam)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ListCommands function
// returns array with all commands from database
func ListCommands() (commands []Exam, err error) {
	result := db.Find(&commands)
	if result.Error != nil {
		return nil, result.Error
	}

	return commands, nil
}

// CreateUser function
func CreateUser(user *User) error {
	// check username and password length
	err = checkInputLength(user.Username, user.Password)
	if err != nil {
		return err
	}

	// set creation date
	user.CreatedAt = time.Now()

	// hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	// create user in database
	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
