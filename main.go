package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(500).JSON(fiber.Map{
			"Response": "Hello World",
		})
	})

	app.Listen("localhost:8000")
}
