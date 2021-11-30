package routes

import (
	"github.com/ertugrul-k/goap/utility"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(r *mux.Router, db *mongo.Database) {
	InitUserRoutes(r, db)
	utility.RouteWalk(r)
}
