package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty" binding:"required"`
	Name      string             `json:"name" bson:"name"`
	Surname   string             `json:"surname" bson:"surname"`
	PinCode   string             `json:"pincode" bson:"pincode"`
	Room      primitive.ObjectID `json:"room" bson:"room"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// type GobotTime time.Time

// func (t GobotTime) MarshalJSON() ([]byte, error) {
// 	//do your serializing here
// 	CreatedAt := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
// 	return []byte(stamp), nil
// }
