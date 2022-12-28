package jwtauth

import (
	"testing"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	helpers "github.com/abodsakah/BTH-appen/backend/api/src/Helpers"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseJwtAuth(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Log("DEV: Could not load .env file")
	}
	dbP, err := db.SetupDatabase()
	assert.Nil(t, err, "Database can not be connected to")
	helpers.DbGorm = dbP
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestGenerateJWT(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestUser)
  _, err := GenerateJWT(helpers.TestEntryIndex)
	assert.Nil(t, err, "After calling generateJWT with valid user id, it shall not return an error")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestValidateJWT1(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestUser)
  token, _ := GenerateJWT(helpers.TestEntryIndex)
  _, err := ValidateJWT(token)
	assert.Nil(t, err, "After calling validateJWT with valid token, it shall not return an error")
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestValidateJWT2(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestUser)
  token := "incorrect-token"
  _, err := ValidateJWT(token)
	assert.NotNil(t, err, "After calling validateJWT with valid token, it shall not return an error")
	_ = helpers.FixtureWrapNonCreate(t)
}
