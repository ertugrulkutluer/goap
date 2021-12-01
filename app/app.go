package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ertugrul-k/goap/db"
	. "github.com/ertugrul-k/goap/db"
	"github.com/ertugrul-k/goap/middleware"
	"github.com/ertugrul-k/goap/routes"
	"github.com/ertugrul-k/goap/utility"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	Env  string
	port int
	r    *mux.Router
)

func initialize_db() {
	db.GetDbContext(Env)
}

func createRouter() {
	r = mux.NewRouter()
	r = r.PathPrefix("/api").Subrouter()
	r.Use(middleware.CORSMiddleware)
}

func init() {
	flag.StringVar(&Env, "env", "development", "current env")
	flag.Parse()
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
	fmt.Println("Server Running on localhost:" + port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CompressHandler(r))
	if err != nil {
		log.Fatal(err)
	}
}
