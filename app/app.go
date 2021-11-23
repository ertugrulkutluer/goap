package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ertugrul-k/apgo/middleware"
	"github.com/ertugrul-k/goap/routes"
	"github.com/ertugrul-k/goap/utility"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Serve() {
	mongo_connection_string := utility.GoDotEnvVariable("MONGO_URI")
	port := utility.GoDotEnvVariable("PORT")

	// Mongo connection

	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_connection_string))
	if err != nil {
		log.Fatal(err)
	}

	// Starting server
	if port == "" {
		port = "8080"
		fmt.Println("Defaulting to port %s", port)
	}
	fmt.Println("Listening on port %s", port)
	uri := ":" + port

	router := gin.New()
	v1 := router.Group("/api/v1/")
	r := v1.Use(middleware.CORSMiddleware(), middleware.JSONMiddleware())
	r.GET("/2", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	routes.InitUserRoutes(r)
	srv := &http.Server{
		Addr:    uri,
		Handler: router,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown: ", err)
	}
	fmt.Println("Server exiting")
}
