package db

import (
	"fmt"
	"jchhay/go-rest-api-gin/config"

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

func NewGormClient() {
	c := config.GetConfig()

	if c.Database.Driver == "sqlite" {
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			fmt.Println("db err: ", err)
			panic("failed to connect database")
		}
	}

	sqlDB, err := GetDB().DB()
	if err != nil {
		fmt.Println("db err: ", err)
		panic("failed to retrieve database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(1000)

	// Migrate the schema
	err = DB.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println("db err: ", err)
		panic("failed to migrate database")
	}

	// Create Dummy Data
	DB.Create(&Book{Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2})
	DB.Create(&Book{Title: "Ulysses", Author: "James Joyce", Quantity: 5})
	DB.Create(&Book{Title: "Don Quixote", Author: "Miguel de Cervantes", Quantity: 10})
	DB.Create(&Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 8})
	DB.Create(&Book{Title: "The Odyssey", Author: "Homer", Quantity: 3})

}

func GetDB() *gorm.DB {
	return DB
}
