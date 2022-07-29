package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"

	"github.com/gofiber/fiber/v2"
)

func AddUserHandler(c *fiber.Ctx) {
	user := new(request.AddUserRequest) //take pattern data submission

	if err := c.BodyParser(user); err != nil {
		return err
	}

	// data form pattern submit to pattern entity db user
	userData := entity.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	err := database.DB.Create(&userData).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to Insert Data",
		})
	}
}
