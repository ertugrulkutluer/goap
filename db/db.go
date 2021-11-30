package db

import (
	"context"
	"log"
	"time"

	"github.com/ertugrul-k/goap/utility"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDbContext() (*mongo.Client, context.Context) {
	return connect(
		utility.GoDotEnvVariable("MONGO_URI"),
	)
}

// Establish a connection to database
func connect(mongo_connection_string string) (*mongo.Client, context.Context) {
	// Mongo connection
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_connection_string))
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx

}
