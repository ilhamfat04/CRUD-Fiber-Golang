package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	repointerface "go-fiber/repo-interface"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(500).JSON(fiber.Map{
		"Response": "Hello World using Golang",
	})
}

func AddUserHandler(c *fiber.Ctx) error {
	userRepository := repointerface.RepositoryUser(database.DB)
	user := new(request.AddUserRequest) //take pattern data submission
	if err := c.BodyParser(user); err != nil {
		return err
	}
	// data form pattern submit to pattern entity db user
	userData := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	// err := database.DB.Create(&userData).Error
	userData, err := userRepository.AddBook(userData)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"Message": "Failed to Insert Data",
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Success Insert Data",
		"Data":    userData,
	})
}
