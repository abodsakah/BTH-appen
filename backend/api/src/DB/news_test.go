package db

import (
	"testing"

	fixture "github.com/abodsakah/BTH-appen/backend/api/src/Fixture"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

func TestDatabaseNews(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := SetupDatabase()
	fixture.AssertNoError(t, err, "Database can not be connected to")
	db = dbP
}

func TestCreateNews1(t *testing.T) {
	_, err := fixtureWrap(t, &testNews)
	fixture.AssertNoError(t, err, "After calling create on news there shall be no errors returned")
}

func TestCreateNews2(t *testing.T) {
	fixtureWrap(t, &testNews)
	err := CreateNews(db, testNews)
	fixture.AssertError(t, err, "After calling create on news when it's a duplicate it shall return an error")
}

func TestDeleteNews1(t *testing.T) {
	newsID, _ := fixtureWrap(t, &testNews)
	DeleteNews(db, newsID)
	var news News
	res, _ := fixture.CheckIfDeleted(db, newsID, &news)
	assert.Equal(t, true, res, "When calling delete the article shall be soft-deleted from the database and not come up in a where statement")
}

func TestGetNews1(t *testing.T) {
	fixtureWrap(t, &testNews)
	res, _ := GetNews(db)
	assert.Less(t, 0, len(res), "When calling getNews after a test entry has been created the function call return an array of larger than 0 in size")
}

func TestGetNews2(t *testing.T) {
	newsID, _ := fixtureWrap(t, &testNews)
	DeleteNews(db, newsID)
	res, _ := GetNews(db)
	assert.Equal(t, 0, len(res), "When calling getNews with no entries it shall return an array of size 0")
}
