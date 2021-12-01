package routes

import (
	"github.com/ertugrul-k/goap/core/service"
	"github.com/gorilla/mux"
)

func InitUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", service.FindAll).Methods("GET")
	r.HandleFunc("/users/{_id}", service.FindOne).Methods("GET")
	r.HandleFunc("/users", service.CreateUser).Methods("POST")
}
