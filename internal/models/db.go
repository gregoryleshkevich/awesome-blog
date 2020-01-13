package models

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// InitDB init the database
func InitDB(dbName, dataSourceName string) {
	var err error
	db, err = gorm.Open(dbName, dataSourceName)
	//defer db.Close()

	if err != nil {
		log.Panic(err)
	}
}
