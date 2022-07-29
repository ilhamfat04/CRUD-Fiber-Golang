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
}
