package db

import (
	"fmt"
	"go_crud/db/model"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error

func DBConfig() {
	Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database :(")
	}
	// Migrate the schema
	db := Database.AutoMigrate(&model.Student{})
	if db != nil {
		fmt.Println("DB error!")
	}
	fmt.Println("connected!")
}
