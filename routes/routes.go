package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r fiber.Router) {
	UserRoute(r)
}
