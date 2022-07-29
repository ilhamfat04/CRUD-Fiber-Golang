package main

import (
	"go-fiber/database"
	"go-fiber/pkg/mysql"
	"go-fiber/repositories"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initil DB
	mysql.DatabaseInit()
	userRepository := repositories.RepositoryUser(mysql.DB)

	// run migration
	database.RunMigration()

	// initial route
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	// group route
	groupRouteApi := app.Group("/api/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	// INITIAL ROUTE
	routes.RouteInit(groupRouteApi, userRepository)

	app.Listen("localhost:8000")
}
