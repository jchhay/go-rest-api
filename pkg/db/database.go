package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type Book struct {
	gorm.Model
	Title    string
	Author   string
	Quantity int
}

func NewGormClient(driver string) {
	var db = DB

	if driver == "sqlite" {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			fmt.Println("db err: ", err)
			panic("failed to connect database")
		}
	}

	DB = db

	// Migrate the schema
	err = db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println("db err: ", err)
		panic("failed to migrate database")
	}

	// Create Dummy Data
	db.Create(&Book{Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2})
	db.Create(&Book{Title: "Ulysses", Author: "James Joyce", Quantity: 5})
	db.Create(&Book{Title: "Don Quixote", Author: "Miguel de Cervantes", Quantity: 10})
	db.Create(&Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 8})
	db.Create(&Book{Title: "The Odyssey", Author: "Homer", Quantity: 3})

}

func GetDB() *gorm.DB {
	return DB
}
