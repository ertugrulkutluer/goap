package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ertugrul-k/goap/db"
	"github.com/ertugrul-k/goap/routes"
	"github.com/ertugrul-k/goap/utility"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func initialize() (*mongo.Client, context.Context) {
	client, ctx := db.GetDbContext()
	return client, ctx
}

// Define HTTP request routes
func Serve() {
	client, ctx := initialize()
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	db := client.Database("production")
	routes.InitRoutes(s, db)
	port, err := utility.GoDotEnvVariable("PORT")
	if err != nil {
		port = "8080"
	}
	defer client.Disconnect(ctx)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Running on localhost:" + port)
}
