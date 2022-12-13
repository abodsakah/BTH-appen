package fixture

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// CleanUp
// Fixture clean up of live database, resets everything in given tables and automigrates
// returns err if dropping tables or automigrating fails
func CleanUp(db *gorm.DB, additionalTables []string, tables ...interface{}) error {
	err := db.Migrator().DropTable(tables...)
	if err != nil {
		return err
	}
	for _, strTable := range additionalTables {
		err = db.Migrator().DropTable(strTable)
		if err != nil {
			return err
		}
	}
	err = db.Migrator().AutoMigrate(tables...)
	return err
}

// CheckIfDeleted 
// Checks if entry is deleted 
// Returns err if it true if it doesn't exist, vice versa
func CheckIfDeleted(db *gorm.DB, id uint, table interface{}) (bool, error) {
	err := db.Where("id = ?", id).First(&table).Error
	if err != nil {
		return true, err
	}
	return false, nil
}

func AssertNoError(t *testing.T, err error, message string) {
	assert.Equal(t, nil, err, message)
}

func AssertError(t *testing.T, err error, message string) {
	assert.NotEqual(t, nil, err, message)
}
