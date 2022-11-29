package db

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "gorm.io/gorm"
  // "github.com/abodsakah/BTH-appen/backend/api/src/Fixture"

	"github.com/joho/godotenv"
)

var db *gorm.DB
var tables = []string{"user", "exam", "news", "token"}


func TestDatabase(t *testing.T) {
  err := godotenv.Load("../../../.env")
  assert.NotEqual(t, err, nil, "Database can not be connected to")

  db_p, err := SetupDatabase()
  assert.NotEqual(t, err, nil, "Database can not be connected to")
  db = db_p
}
/*
func TestExample(t *testing.T) {
  fixture.CleanUp(db, tables, &User{}, &Exam{}, &News{}, &Token{})
  assert.Equal(t, 1, 2, "They should be equal")
}
*/
