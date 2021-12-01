package routes

import (
	"net/http"

	"github.com/ertugrul-k/goap/core/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitUserRoutes(r *mux.Router, db *mongo.Database) {
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		service.FindAll(w, r, db)
	}).Methods("GET")
	r.HandleFunc("/users/{_id}", func(w http.ResponseWriter, r *http.Request) {
		service.FindOne(w, r, db)
	}).Methods("GET")
}

// func userCreate(r *mux.Router) {
// 	fmt.Println("test1")
// }
