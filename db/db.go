package db

import (
	"context"
	"log"
	"time"

	c "github.com/ertugrul-k/goap/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	Database *mongo.Database
	Ctx      context.Context
	Client   *mongo.Client
}

var DB database

func GetDbContext(env string) {
	conn_str := c.Config.Database
	connect(conn_str, env)
}

// Establish a connection to database
func connect(mongo_connection_string, env string) {
	// Mongo connection
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_connection_string))
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	DB.Client = client
	DB.Ctx = ctx
	DB.Database = client.Database(c.Config.Env) // client.Database(env)
	err = DB.Client.Connect(DB.Ctx)
	if err != nil {
		log.Fatal(err)
	}
}
