// Package fixture provides fixture
package fixture

import (
	"gorm.io/gorm"
)

// CleanUp function
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
