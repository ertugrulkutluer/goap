package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ertugrul-k/goap/db"
	. "github.com/ertugrul-k/goap/db"
	"github.com/ertugrul-k/goap/middleware"
	"github.com/ertugrul-k/goap/routes"
	"github.com/ertugrul-k/goap/utility"
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
	// flag.StringVar(&Env, "env", "development", "current env")
	// flag.Parse()
	if Env = os.Getenv("env"); Env == "" {
		Env = "staging"
	}
	log.Println(fmt.Sprintf("Working Environment: %s\n", Env))
}

// Define HTTP request routes
func Serve() {
	initialize_db()
	createRouter()
	routes.InitRoutes(r)
	port, err := utility.GoDotEnvVariable("PORT", Env)
	if err != nil {
		port = "8080"
	}

	defer DB.Client.Disconnect(DB.Ctx)

	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server Running on localhost:" + port)
	}
}
