package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coin struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty" binding:"required"`
	Name      string             `bson:"name" json:"name"`
	Code      string             `bson:"code" json:"code"`
	Order     int                `bson:"order" json:"order"`
	Tags      []string           `bson:"tags" json:"tags"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// type GobotTime time.Time

// func (t GobotTime) MarshalJSON() ([]byte, error) {
// 	//do your serializing here
// 	CreatedAt := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
// 	return []byte(stamp), nil
// }
