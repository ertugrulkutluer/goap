package repo

import (
	"context"
	"log"
	"math"
	"strconv"

	"github.com/ertugrul-k/goap/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func FindAll(ctx context.Context, coll *mongo.Collection, r *fiber.Ctx) error {
	var users []models.User
	filter := bson.M{}
	findOptions := options.Find()

	if s := r.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"email": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"name": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page, _ := strconv.Atoi(r.Query("page", "1"))
	limitVal, _ := strconv.Atoi(r.Query("limit", "10"))
	var limit int64 = int64(limitVal)

	total, _ := coll.CountDocuments(ctx, bson.M{})

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := coll.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Users Not found",
			"error":   err,
		})
	}
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return r.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      users,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})

}

func FindOne(ctx context.Context, coll *mongo.Collection, r *fiber.Ctx) error {
	var user models.User
	user_id, _ := primitive.ObjectIDFromHex(r.Params("_id"))
	filter := bson.M{"_id": user_id}
	result := coll.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found.",
			"error":   err,
		})
	}
	err := result.Decode(&user)
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found.",
			"error":   err,
		})
	}
	return r.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})
}

func Create(ctx context.Context, coll *mongo.Collection, r *fiber.Ctx) error {
	var user *models.User
	if err := r.BodyParser(&user); err != nil {
		log.Println(err)
		return r.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body.",
			"error":   err,
		})
	}
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			log.Fatal("Email already exists in database.")
		default:
			log.Fatal("Error while inserting data.")
		}
	}
	return r.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "User inserted successfully.",
	})
}

func Update(ctx context.Context, coll *mongo.Collection, r *fiber.Ctx) error {
	var user *models.User
	if err := r.BodyParser(&user); err != nil {
		log.Println(err)
		return r.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body.",
			"error":   err,
		})
	}
	obj_id, err := primitive.ObjectIDFromHex(r.Params("_id"))
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}

	update := bson.M{
		"$set": user,
	}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": obj_id}, update)
	if err != nil {
		return r.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update.",
			"error":   err.Error(),
		})
	}
	return r.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User uptated successfully.",
	})
}

func Delete(ctx context.Context, coll *mongo.Collection, r *fiber.Ctx) error {
	obj_id, err := primitive.ObjectIDFromHex(r.Params("_id"))
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}
	_, err = coll.DeleteOne(ctx, bson.M{"_id": obj_id})
	if err != nil {
		return r.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User failed to update.",
			"error":   err.Error(),
		})
	}
	return r.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully.",
	})
}
