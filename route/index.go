package route

import "github.com/gofiber/fiber/v2"

func RouteInit(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(500).JSON(fiber.Map{
			"Response": "Hello World",
		})
	})

}
