package db

import (
	"testing"

	helpers "github.com/abodsakah/BTH-appen/backend/api/src/Helpers"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseUser(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Log("DEV: Could not load .env file")
	}
	dbP, err := SetupDatabase()
	assert.Nil(t, err, "Database can not be connected to")
	helpers.DbGorm = dbP
}

func TestUserCreate1(t *testing.T) {
	_, err := helpers.FixtureWrapCreate(t, helpers.TestUser)
	assert.Nil(t, err, "When calling create, with no duplicates, it shall return no errors")
}

func TestUserCreate2(t *testing.T) {
	_, _ = helpers.FixtureWrapCreate(t, helpers.TestUser)
	temp := *helpers.TestUser
	err := CreateUser(helpers.DbGorm, &temp)
	assert.NotNil(t, err, "When calling create, with duplicates, it shall return errors")
}

func TestUserIsRole1(t *testing.T) {
	id, _ := helpers.FixtureWrapCreate(t, helpers.TestUser)
	res, _ := IsRole(helpers.DbGorm, id, "student")
	assert.Equal(t, true, res, "User with role \"student\" shall make the function return true when its given \"student\"")
}

func TestUserIsRole2(t *testing.T) {
	id, _ := helpers.FixtureWrapCreate(t, helpers.TestUser)
	res, _ := IsRole(helpers.DbGorm, id, "admin")
	assert.NotEqual(t, true, res, "User with role \"student\" shall make function return false when its given admin")
}

func TestUserAuth1(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	temp := *helpers.TestUser
	_ = CreateUser(helpers.DbGorm, &temp)
	_, err := AuthUser(helpers.DbGorm, helpers.TestUser.Username, helpers.TestUser.Password)
	assert.Nil(t, err, "Authenticating created user with correct information shall return no errors")
}

func TestUserAuth2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	temp := *helpers.TestUser
	_ = CreateUser(helpers.DbGorm, &temp)
	_, err := AuthUser(helpers.DbGorm, helpers.TestUser.Username, "IncorrectPassword")
	assert.NotNil(t, err, "Authenticating created user with incorrect information shall return errors")
}
