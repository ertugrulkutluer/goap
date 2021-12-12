package service

import (
	"context"
	"time"

	"github.com/ertugrul-k/goap/core/repo"
	. "github.com/ertugrul-k/goap/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func define() (context.Context, *mongo.Collection) {
	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	return ctx, coll

}

func FindAll(r *fiber.Ctx) error {
	cx, cl := define()
	return repo.FindAll(cx, cl, r)
}

func FindOne(r *fiber.Ctx) error {
	cx, cl := define()
	return repo.FindOne(cx, cl, r)
}

func Create(r *fiber.Ctx) error {
	cx, cl := define()
	return repo.Create(cx, cl, r)
}

func Update(r *fiber.Ctx) error {
	cx, cl := define()
	return repo.Update(cx, cl, r)
}

func Delete(r *fiber.Ctx) error {
	cx, cl := define()
	return repo.Delete(cx, cl, r)
}

func Login(r *fiber.Ctx) error {
	cx, cl := define()
	return repo.Login(cx, cl, r)
}

func Register(r *fiber.Ctx) error {
	cx, cl := define()
	return repo.Register(cx, cl, r)
}
