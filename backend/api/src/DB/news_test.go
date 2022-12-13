package db

import (
	"testing"

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
	assert.Nil(t, err, "Database can not be connected to")
	db = dbP
}

func TestCreateNews1(t *testing.T) {
	_, err := fixtureWrap(t, &testNews)
	assert.Nil(t, err, "After calling create, with no duplicates, no errors shall be returned")
}

func TestCreateNews2(t *testing.T) {
	_, _ = fixtureWrap(t, &testNews)
	err := CreateNews(db, testNews)
	assert.NotNil(t, err, "After calling create, with duplicates, errors shall be returned")
}

func TestDeleteNews1(t *testing.T) {
	newsID, _ := fixtureWrap(t, &testNews)
	_, err := DeleteNews(db, newsID)
	assert.Nil(t, err, "After calling delete on news, it should not return any errors")
}

func TestGetNews1(t *testing.T) {
	_, _ = fixtureWrap(t, &testNews)
	res, _ := GetNews(db)
	assert.Less(t, 0, len(res), "When calling getNews after a test entry has been created the function call return an array of larger than 0 in size")
}

func TestGetNews2(t *testing.T) {
	newsID, _ := fixtureWrap(t, &testNews)
	_, _ = DeleteNews(db, newsID)
	res, _ := GetNews(db)
	assert.Equal(t, 0, len(res), "When calling getNews with no entries it shall return an array of size 0")
}
