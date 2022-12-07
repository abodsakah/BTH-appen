package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var (
	db               *gorm.DB
	additionalTables = []string{"exam_users"}
)

var testUser = &User{
	Name:     "Test Testsson",
	Username: "test",
	Password: "pass",
	Role:     "student",
}

func createUserWrap() error {
	fmt.Printf("Username: %s\n", testUser.Username)
	fmt.Printf("Oassword: %s\n", testUser.Password)
	temp := *testUser
	err := CreateUser(db, testUser)
	*testUser = temp
	if err != nil {
		return err
	}
	return nil
}

func fixtureWrap(t *testing.T) error {
	err := cleanUp(db, additionalTables)
	if err != nil {
		t.Fatal(err)
	}
	createUserWrap()
	return err
}

func assertNoError(t *testing.T, err error, message string) {
	assert.Equal(t, nil, err, message)
}

func assertError(t *testing.T, err error, message string) {
	assert.NotEqual(t, nil, err, message)
}

func TestDatabase(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := SetupDatabase()
	assertNoError(t, err, "Database can not be connected to")
	db = dbP
}

// Tries to create user "Test" when there is none in the database
func TestUserCreate1(t *testing.T) {
	err := fixtureWrap(t)
	assertNoError(t, err, "API endpoint should be able to create user if no duplicates are present")
}

// Tries to create user "Test" when there already is in database
func TestUserCreate2(t *testing.T) {
	fixtureWrap(t)
	err := createUserWrap()
	assertError(t, err, "API should return error as duplicate already exists")
}

// Tries to create exam when there is none in database
func TestUserIsRole1(t *testing.T) {
	fixtureWrap(t)
	user, _ := GetUserByName(db, testUser.Username)
	res, _ := IsRole(db, user.ID, "student")
	assert.Equal(t, true, res, "User with role \"student\" shall make the function return true when its given \"student\"")
}

func TestUserIsRole2(t *testing.T) {
	fixtureWrap(t)
	user, _ := GetUserByName(db, testUser.Username)
	res, _ := IsRole(db, user.ID, "admin")
	assert.NotEqual(t, true, res, "User with role \"student\" shall make function return false when its given admin")
}

func TestUserAuth1(t *testing.T) {
	fixtureWrap(t)
	_, err := AuthUser(db, testUser.Username, testUser.Password)
	assertNoError(t, err, "Authenticating created user with correct information shall return no errors")
}

func TestUserAuth2(t *testing.T) {
	fixtureWrap(t)
	_, err := AuthUser(db, testUser.Username, "IncorrectPassword")
	assertError(t, err, "Authenticating created user with correct information shall return no errors")
}
