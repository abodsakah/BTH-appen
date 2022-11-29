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
		err = db.Migrator().DropTable(strTable)
		if err != nil {
			return err
		}
	}
	err = db.Migrator().AutoMigrate(&User{}, &Exam{}, &News{}, &Token{})
	return err
}
