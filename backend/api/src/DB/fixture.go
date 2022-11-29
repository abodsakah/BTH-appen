package db

import (
	"gorm.io/gorm"
)

func cleanUp(db *gorm.DB, additionalTables []string) error {
	err := db.Migrator().DropTable(&User{}, &Exam{}, &News{}, &Token{})
	if err != nil {
		return err
	}
	for _, strTable := range additionalTables {
		db.Migrator().DropTable(strTable)
	}
	err = db.Migrator().AutoMigrate(&User{}, &Exam{}, &News{}, &Token{})
	return err
}
