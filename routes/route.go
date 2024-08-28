package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhamadijlal/gobasic/controllers"
)

func RouterApp(c *fiber.App) {
	c.Get("/api/allUsers", controllers.UserControllerShow)
	c.Get("/api/user/:id", controllers.UserControllerGetByID)
	c.Patch("/api/user/:id", controllers.UserControllerUpdate)
	c.Post("/api/create", controllers.UserControllerCreate)
	c.Delete("/api/user/:id", controllers.UserControllerDelete)
}
