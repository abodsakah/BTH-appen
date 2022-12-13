package db

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseUser(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := SetupDatabase()
  assert.Nil(t, err, "Database can not be connected to")
	db = dbP
}

func TestUserCreate1(t *testing.T) {
	_, err := fixtureWrapUser(t)
  assert.Nil(t, err, "When calling create, with no duplicates, it shall return no errors")
}

func TestUserCreate2(t *testing.T) {
	_, _ = fixtureWrapUser(t)
	_, err := createUserWrap()
  assert.NotNil(t, err, "When calling create, with duplicates, it shall return errors")
}

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
	_, _ = fixtureWrapUser(t)
	_, err := AuthUser(db, testUser.Username, testUser.Password)
  assert.Nil(t, err, "Authenticating created user with correct information shall return no errors")
}

func TestUserAuth2(t *testing.T) {
	_, _ = fixtureWrapUser(t)
	_, err := AuthUser(db, testUser.Username, "IncorrectPassword")
  assert.NotNil(t, err, "Authenticating created user with incorrect information shall return errors")
}
