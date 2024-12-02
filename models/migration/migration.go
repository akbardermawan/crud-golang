package migration

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/models/entity"
	"log"
)

func RunMigration() {
	// Pastikan koneksi database tidak nil
	if database.DB == nil {
		log.Panicln("Database connection is not initialized")
	}

	// Lakukan migrasi untuk model User
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Panicln("Migration failed:", err)
	}
	fmt.Println("Database successfully migrated")
}