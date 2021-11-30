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
	port := utility.GoDotEnvVariable("PORT")
	defer client.Disconnect(ctx)
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Running on localhost:" + port)
}
