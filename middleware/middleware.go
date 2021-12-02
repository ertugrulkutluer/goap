package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware(r *fiber.App) fiber.Router {
	return r.Use(cors.New())
	// return r.Use(cors.New(cors.Config{
	// 	AllowOrigins: "https://gofiber.io, https://gofiber.net",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))
}
