package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty" binding:"required"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	Surname   string             `json:"surname" bson:"surname" validate:"required"`
	PinCode   string             `json:"pincode" bson:"pincode"`
	Room      primitive.ObjectID `json:"room" bson:"room" default:"null"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	IsActive  *bool              `json:"is_active" bson:"is_active"`
	Email     string             `json:"email" bson:"email" validate:"required,email,min=6,max=32"`
	Role      string             `json:"role" bson:"role" validate:"required"`
}
