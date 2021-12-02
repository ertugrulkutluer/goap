package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(r fiber.Router) {
	InitUserRoutes(r)
	// utility.RouteWalk(r)
}
