package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ertugrul-k/goap/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongo_connection_string := goDotEnvVariable("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_connection_string))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database("production")
	coins_collection := db.Collection("coins")
	coin := models.Coin{
		Name:      "Bitcoin",
		Code:      "BTC",
		Order:     1,
		CreatedAt: time.Now()}
	result, err := coins_collection.InsertOne(ctx, coin)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": dotenv,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
