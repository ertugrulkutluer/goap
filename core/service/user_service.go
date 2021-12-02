package service

import (
	"context"
	"time"

	"github.com/ertugrul-k/goap/core/repo"
	. "github.com/ertugrul-k/goap/db"
	"github.com/gofiber/fiber/v2"
)

func FindAll(r *fiber.Ctx) error {
	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := repo.FindAll(ctx, coll, r)
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Users Not found",
			"error":   err,
		})
	}
	return err
}

func FindOne(r *fiber.Ctx) error {
	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := repo.FindOne(ctx, coll, r)
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	return err
}

func Create(r *fiber.Ctx) error {
	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := repo.Create(ctx, coll, r)
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	return err
}

func Update(r *fiber.Ctx) error {
	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := repo.Update(ctx, coll, r)
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	return err
}

func Delete(r *fiber.Ctx) error {
	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := repo.Delete(ctx, coll, r)
	if err != nil {
		return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User Not found",
			"error":   err,
		})
	}
	return err
}
