package controllers

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhamadijlal/gobasic/database"
	"github.com/muhamadijlal/gobasic/models/entity"
	"github.com/muhamadijlal/gobasic/models/request"
)

func UserControllerShow(c *fiber.Ctx) error {
	var users []entity.User

	err := database.DB.Find(&users).Error

	if err != nil {
		log.Println(err)
	}

	return c.JSON(users)
}

func UserControllerCreate(c *fiber.Ctx) error {
	user := new(request.UserRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed create new user",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		Name:  user.Name,
		Email: user.Email,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create new user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success create new user",
		"data":    newUser,
	})
}

func UserControllerGetByID(c *fiber.Ctx) error {
	var user []entity.User
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Null Params",
		})
		return nil
	}
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Params",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User data found",
		"data":    user,
	})
}

func UserControllerUpdate(c *fiber.Ctx) error {
	userUpdate := new(request.UserRequest)
	id := c.Params("id")

	if err := c.BodyParser(userUpdate); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed update user",
			"error":   err.Error(),
		})
	}

	var user entity.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Params",
		})
		return nil
	}

	user.Name = userUpdate.Name
	user.Email = userUpdate.Email

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	// Return a success response
	return c.JSON(fiber.Map{
		"message": "Successfully updated user",
	})

}

func UserControllerDelete(c *fiber.Ctx) error {

	id := c.Params("id")
	var user entity.User

	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
		return nil
	}
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to find user",
			"error":   err.Error(),
		})
		return nil
	}

	// Delete the user from the database
	if err := database.DB.Delete(&user).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
		return nil
	}

	// Return a success response
	return c.JSON(fiber.Map{
		"message": "Successfully deleted user",
	})
}
