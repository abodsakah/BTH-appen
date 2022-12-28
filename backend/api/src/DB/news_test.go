package db

import (
	"testing"

	helpers "github.com/abodsakah/BTH-appen/backend/api/src/Helpers"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseNews(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Log("DEV: Could not load .env file")
	}
	dbP, err := SetupDatabase()
	assert.Nil(t, err, "Database can not be connected to")
	helpers.DbGorm = dbP
	_ = helpers.FixtureWrapNonCreate(t)
}

func TestCreateNews1(t *testing.T) {
	err := helpers.FixtureWrapCreate(t, &helpers.TestNews)
	assert.Nil(t, err, "After calling create, with no duplicates, no errors shall be returned")
}

func TestCreateNews2(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestNews)
	err := CreateNews(helpers.DbGorm, helpers.TestNews)
	assert.NotNil(t, err, "After calling create, with duplicates, errors shall be returned")
}

func TestDeleteNews1(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestNews)
	_, err := DeleteNews(helpers.DbGorm, helpers.TestEntryIndex)
	assert.Nil(t, err, "After calling delete on news, it should not return any errors")
}

func TestGetNews1(t *testing.T) {
	_ = helpers.FixtureWrapCreate(t, &helpers.TestNews)
	res, _ := GetNews(helpers.DbGorm)
	assert.Less(t, 0, len(res), "When calling getNews after a test entry has been created the function call return an array of larger than 0 in size")
}

func TestGetNews2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	_, _ = DeleteNews(helpers.DbGorm, helpers.TestEntryIndex)
	res, _ := GetNews(helpers.DbGorm)
	assert.Equal(t, 0, len(res), "When calling getNews with no entries it shall return an array of size 0")
}
