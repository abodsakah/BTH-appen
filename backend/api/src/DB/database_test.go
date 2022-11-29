package db

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "gorm.io/gorm"

	"github.com/joho/godotenv"
)

var db *gorm.DB
var additionalTables = []string{"exam_users"}

func TestDatabase(t *testing.T) {
  err := godotenv.Load("../../../.env")
  assert.Equal(t, nil, err, "Database can not be connected to")

  db_p, err := SetupDatabase()
  assert.NotEqual(t, nil, err, "Database can not be connected to")
  db = db_p
}
func TestExample(t *testing.T) {
  cleanUp(db, additionalTables)
  assert.Equal(t, 1, 2, "They should be equal")
  _, err := SetupDatabase()
  assert.Equal(t, nil, err, "Database can not be connected to")
  cleanUp(db, additionalTables)
}

