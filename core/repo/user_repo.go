package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/ertugrul-k/goap/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func FindAll(ctx context.Context, coll *mongo.Collection) []models.Coin {
	var users []models.Coin
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.Coin
		cursor.Decode(&user)
		fmt.Printf("%+v\n", user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func FindOne(ctx context.Context, coll *mongo.Collection, id string) models.Coin {
	// var (
	// 	modelStruct models.Coin
	// )
	user_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": user_id}
	result := coll.FindOne(ctx, filter)
	var user models.Coin
	result.Decode(&user)
	return user
}

// func CreateUser(ctx context.Context, coll *mongo.Collection, body io.ReadCloser) models.Coin {
// 	var user models.Coin
// 	_ = json.NewDecoder(r.Body).Decode(user)
// 	user_id, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.M{"_id": user_id}
// 	result := coll.FindOne(ctx, filter)
// 	var user models.Coin
// 	result.Decode(&user)
// 	return user
// }
