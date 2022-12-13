// Unit tests for route module
package routes

import (
	"testing"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
  dbGorm *gorm.DB
)

func TestDatabaseRoutes(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := db.SetupDatabase()
	assert.Nil(t, err, "Database can not be connected to")
	dbGorm = dbP
}


