package routes

import (
	user_service "github.com/ertugrul-k/goap/core/service"
	"github.com/gofiber/fiber/v2"
)

func InitUserRoutes(r fiber.Router) {

	users := r.Group("/users")

	users.Get("/", user_service.FindAll)
	users.Get("/:_id", user_service.FindOne)
	users.Post("/", user_service.Create)
	users.Put("/:_id", user_service.Update)
	users.Delete("/:_id", user_service.Delete)

	users.Post("/login", user_service.Login)
	users.Post("/register", user_service.Register)
}
