package database

import (
	"log"

	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {

		log.Fatalf("Failed to connect to DB: %v", err)

	}
	database.AutoMigrate(&model.UserEntity{})

	log.Println("In-memory DB connected and migrated!")

	DB = database

}
