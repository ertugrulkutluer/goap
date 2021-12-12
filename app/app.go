package app

import (
	"fmt"
	"log"
	"os"

	c "github.com/ertugrul-k/goap/config"
	"github.com/ertugrul-k/goap/db"
	. "github.com/ertugrul-k/goap/db"
	"github.com/ertugrul-k/goap/middleware"
	"github.com/ertugrul-k/goap/routes"
	"github.com/gofiber/fiber/v2"
)

var (
	Env  string
	port int
	r    fiber.Router
	app  *fiber.App
)

func initialize_db() {
	db.GetDbContext(Env)
}

func createRouter() {
	app = fiber.New()
	app_middleware := middleware.CORSMiddleware(app)
	r = app_middleware.Group("/api")
}

func init() {
	if Env = os.Getenv("env"); Env == "" {
		Env = "staging"
	}
	log.Println(fmt.Sprintf("Working Environment: %s\n", Env))
	c.Config.Read(Env)
}

// Define HTTP request routes
func Serve() {
	initialize_db()
	createRouter()
	routes.InitRoutes(r)
	port := c.Config.Port
	defer DB.Client.Disconnect(DB.Ctx)

	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server Running on localhost:" + port)
	}
}
