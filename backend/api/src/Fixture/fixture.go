package fixture

import (

	"gorm.io/gorm"
)

func CleanUp(db *gorm.DB, tables []string, models ...interface{}) error {
  var err error
  for i := 0;i < len(tables);i++ {  
    err = db.Migrator().DropTable(tables[i])
    if err != nil {
      return err
    }
  }
  err = db.Migrator().AutoMigrate(models)
  return err
} 

