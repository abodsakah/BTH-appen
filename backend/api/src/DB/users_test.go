package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	fixture "github.com/abodsakah/BTH-appen/backend/api/src/Fixture"
	"github.com/joho/godotenv"
)

func TestDatabaseUser(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := SetupDatabase()
	fixture.AssertNoError(t, err, "Database can not be connected to")
	db = dbP
}

// Tries to create user "Test" when there is none in the database
func TestUserCreate1(t *testing.T) {
	_, err := fixtureWrapUser(t)
	fixture.AssertNoError(t, err, "API endpoint should be able to create user if no duplicates are present")
}

// Tries to create user "Test" when there already is in database
func TestUserCreate2(t *testing.T) {
	fixtureWrapUser(t)
	_, err := createUserWrap()
	fixture.AssertError(t, err, "API should return error as duplicate already exists")
}

// Tries to create exam when there is none in database
func TestUserIsRole1(t *testing.T) {
  id, _ := fixtureWrapUser(t)
	res, _ := IsRole(db, id, "student")
	assert.Equal(t, true, res, "User with role \"student\" shall make the function return true when its given \"student\"")
}

func TestUserIsRole2(t *testing.T) {
  id, _ := fixtureWrapUser(t)
	res, _ := IsRole(db, id, "admin")
	assert.NotEqual(t, true, res, "User with role \"student\" shall make function return false when its given admin")
}

func TestUserAuth1(t *testing.T) {
	fixtureWrapUser(t)
	_, err := AuthUser(db, testUser.Username, testUser.Password)
	fixture.AssertNoError(t, err, "Authenticating created user with correct information shall return no errors")
}

func TestUserAuth2(t *testing.T) {
	fixtureWrapUser(t)
	_, err := AuthUser(db, testUser.Username, "IncorrectPassword")
	fixture.AssertError(t, err, "Authenticating created user with correct information shall return no errors")
}
