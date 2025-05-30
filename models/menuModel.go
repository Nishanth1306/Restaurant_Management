package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_Date *time.Time         `json:"start_Date"`
	End_Date   *time.Time         `json:"end_Date"`
	Created_at *time.Time         `json:"created_Date"`
	Updated_at *time.Time         `json:"updated_Date"`
	Menu_id    string             `json:"food_id"`
}
