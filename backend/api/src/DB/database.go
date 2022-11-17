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

	// FIX: All tests/experiments should be removed from here eventually.
	//**************************
	// Some test DB operations.
	//**************************

	// create admin account
	user := &User{Username: "admin", Password: "pass"}
	err = CreateUser(db, user)
	if err != nil {
		log.Println(err)
	}

	// get first user
	var userOne User
	db.First(&userOne)

	// create some exams
	err = CreateExam(db, &Exam{
		CourseCode: "DV1337",
		StartDate:  time.Now(),
		Users:      []*User{&userOne},
	})
	if err != nil {
		log.Println(err)
	}
	err = CreateExam(db, &Exam{
		CourseCode: "MA6666",
		StartDate:  time.Now().Add(-(time.Hour * 2)),
		Users:      []*User{&userOne},
	})
	if err != nil {
		log.Println(err)
	}
	err = CreateExam(db, &Exam{
		CourseCode: "PA6969",
		StartDate:  time.Now().AddDate(0, 0, -1),
		Users:      []*User{&userOne},
	})
	if err != nil {
		log.Println(err)
	}

	// create user account
	err = CreateUser(db, &User{Username: "user", Password: "pass"})
	if err != nil {
		log.Println(err)
	}
	// auth user and get userID
	userID, err := AuthUser(db, "user", "pass")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("userID: ", userID)
	}
	err = RegisterToExam(db, 1, userID)
	if err != nil {
		log.Println(err)
	}
	err = UnregisterFromExam(db, 1, userOne.ID)
	if err != nil {
		log.Println(err)
	}

	// list exams
	exams, err := ListExams(db)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("\nExams: %#v\n", exams)

	// get exams with users preloaded
	// and print each exam and it's users
	fmt.Println("\n------ exams with registered users preloaded -------")
	var examsPreloaded []Exam
	err = db.Model(&Exam{}).Preload("Users").Find(&examsPreloaded).Error
	if err != nil {
		log.Println(err)
	} else {
		for key, obj := range examsPreloaded {
			fmt.Printf("--------- Exam %d ------------\n", key)
			fmt.Printf("%#v\n", obj)
			fmt.Printf("\nExam %d Users: \n", key)
			for _, user := range obj.Users {
				fmt.Printf("%#v\n", *user)
			}
			fmt.Printf("------------------------------\n")
		}
	}

	// get users with registered exams preloaded
	// and print each exam and it's users
	fmt.Println("\n------ users with registered exams preloaded -------")
	var usersPreloaded []User
	err = db.Model(&User{}).Preload("Exams").Find(&usersPreloaded).Error
	if err != nil {
		log.Println(err)
	} else {
		for key, obj := range usersPreloaded {
			fmt.Printf("--------- User %d ------------\n", key)
			fmt.Printf("%#v\n", obj)
			fmt.Printf("\nUser %d Exams: \n", key)
			for _, exam := range obj.Exams {
				fmt.Printf("%#v\n", *exam)
			}
			fmt.Printf("------------------------------\n")
		}
	}

	//**************************
	// End test DB operations.
	//**************************

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
