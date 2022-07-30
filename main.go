package main

import (
	"go-fiber/database"
	"go-fiber/pkg/mysql"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	// initial route
	app := fiber.New()
	// app := mux.NewRouter()
	// app.HandleFunc("/", YourHandler)

	// group route
	groupRouteApi := app.Group("/api/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	// initial route
	routes.RouteInit(groupRouteApi)

	app.Listen("localhost:8000")
	// http.ListenAndServe("localhost:8080", app)
}
