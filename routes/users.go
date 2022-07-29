package routes

import (
	"go-fiber/handlers"
	"go-fiber/repositories"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r fiber.Router, userRepository repositories.UserRepository) {
	h := handlers.HandlerUser(userRepository)

	r.Get("/", h.HelloWorld)
	r.Post("/user", h.CreateUser)
	r.Get("/users", h.FindUsers)
	r.Get("/user/:id", h.GetUser)
	r.Patch("/user/:id", h.UpdateUser)
	r.Delete("/user/:id", h.DeleteUser)
}
