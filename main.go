package main

import (
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	// GROUP ROUTE
	api := app.Group("/api/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	// INITIAL ROUTE
	route.RouteInit(api)

	app.Listen("localhost:8000")
}
