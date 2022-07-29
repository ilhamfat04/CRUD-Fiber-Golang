package handler

import (
	"go-fiber/model/entity"
	"go-fiber/model/request"

	"go-fiber/repointerface"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	UserRepository repointerface.Repository
}

func HandlerUser(UserRepository repointerface.Repository) *userHandler {
	return &userHandler{UserRepository} // agar bisa diakses di main
}

func (h *userHandler) HelloWorld(c *fiber.Ctx) error {
	return c.Status(500).JSON(fiber.Map{
		"Response": "Hello World using Golang",
	})
}

func (h *userHandler) AddUserHandler(c *fiber.Ctx) error {
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
	userDataInput, err := h.UserRepository.AddUser(userData)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"Message": "Failed to Insert Data",
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Success Insert Data",
		"Data":    userDataInput,
	})
}
