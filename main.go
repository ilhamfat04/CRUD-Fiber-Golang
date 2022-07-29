package main

import (
	"go-fiber/database"
	"go-fiber/database/migration"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DB
	database.DatabaseInit()

	// RUN MIGRATION
	migration.RunMigration()

	// INITIAL ROUTE
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
