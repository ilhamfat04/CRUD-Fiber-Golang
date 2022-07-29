package route

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r fiber.Router) {
	r.Get("/", handler.HelloWorld)
	r.Post("/user", handler.AddUserHandler)
}
