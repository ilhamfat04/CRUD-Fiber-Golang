package routes

import (
	"go-fiber/handlers"
	"go-fiber/pkg/mysql"
	"go-fiber/repositories"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.Get("/", h.HelloWorld)
	r.Post("/user", h.CreateUser)
	r.Get("/users", h.FindUsers)
	r.Get("/user/:id", h.GetUser)
	r.Patch("/user/:id", h.UpdateUser)
	r.Delete("/user/:id", h.DeleteUser)
}
