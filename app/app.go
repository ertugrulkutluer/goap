package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
		logging.Info("Defaulting to port %s", port)
	}
	logging.Info("Listening on port %s", port)
	uri := ":" + port
	a.onStartup()

	router := gin.New()
}
