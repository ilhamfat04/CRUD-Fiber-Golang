package route

import (
	"go-fiber/handler"
	"go-fiber/repointerface"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r fiber.Router, userRepository repointerface.Repository) {
	h := handler.HandlerUser(userRepository)

	r.Get("/", h.HelloWorld)
	r.Post("/user", h.AddUserHandler)
	r.Get("/users", h.GetUsersHandler)
	r.Get("/user/:id", h.GetUserHandler)
	r.Put("/user/:id", h.UpdateUserHandler)
}
