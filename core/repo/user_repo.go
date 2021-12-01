package repo

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/ertugrul-k/goap/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func FindAll(ctx context.Context, coll *mongo.Collection) *[]models.User {
	var users []models.User
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return &users
}

func FindOne(ctx context.Context, coll *mongo.Collection, id string) *models.User {
	// var (
	// 	modelStruct models.Coin
	// )
	user_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": user_id}
	result := coll.FindOne(ctx, filter)
	var user models.User
	result.Decode(&user)
	return &user
}

func CreateUser(ctx context.Context, coll *mongo.Collection, body io.ReadCloser) *models.User {
	var user models.User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			log.Fatal("username or email already exists in database.")
		default:
			log.Fatal("Error while inserting data.")
		}
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user
}
