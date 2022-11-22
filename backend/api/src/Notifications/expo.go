// Package notifications provides notifications
package notifications

import (
	"fmt"

	"github.com/abodsakah/BTH-appen/backend/api/src/DB"
	// expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

// StartServer function
func StartServer(gormObj *gorm.DB) error {
	// setup GORM database object
	gormDB = gormObj

	exams, err := db.GetExamsDueSoon(gormDB)
	if err != nil {
		return err
	}

	fmt.Println("Exams due soon: ", exams)

	return nil
}
