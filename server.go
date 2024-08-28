package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhamadijlal/gobasic/database"
	"github.com/muhamadijlal/gobasic/database/migration"
	"github.com/muhamadijlal/gobasic/routes"
)

func main() {
	database.ConnectDB()
	migration.RunMigrate()

	app := fiber.New()
	routes.RouterApp(app)

	app.Listen(":3000")
}
