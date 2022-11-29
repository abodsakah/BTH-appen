package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var (
	db               *gorm.DB
	additionalTables = []string{"exam_users"}
)

func fixtureWrap(t *testing.T) {
	err := cleanUp(db, additionalTables)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDatabase(t *testing.T) {
	err := godotenv.Load("../../../.env")
	assert.Equal(t, nil, err, "Database can not be connected to")

	dbP, err := SetupDatabase()
	assert.NotEqual(t, nil, err, "Database can not be connected to")
	db = dbP
}

func TestExample(t *testing.T) {
	fixtureWrap(t)
	assert.Equal(t, 1, 2, "They should be equal")
	_, err := SetupDatabase()
	assert.Equal(t, nil, err, "Database can not be connected to")
	fixtureWrap(t)
}
