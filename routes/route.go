package routes

import (
	"github.com/ertugrul-k/goap/utility"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	InitUserRoutes(r)
	utility.RouteWalk(r)
}
