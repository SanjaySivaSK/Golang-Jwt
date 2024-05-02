package database

import (
	model "MyModulw/Model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)

var Db *gorm.DB
var dbError error

func Connect(connectionString string) {
	Db, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Db.AutoMigrate(&model.User{})
	log.Println("Database Migration Completed!")
}
