package main

import (
	"go-fiber/database"
	"go-fiber/models/migration"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initial DATABASE
	database.DatabaseInit()	
	//menjalankan MIgrasi
	migration.RunMigration()

	app := fiber.New()

	// initial route
	routes.RouteInit(app)




	app.Listen(":5000")
	
}