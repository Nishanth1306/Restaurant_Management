package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name"`
	Last_name     *string            `json:"last_name"`
	Password      *string            `json:"Password"`
	Email         *string            `json:"email"`
	Avatar        string             `json:"avatar"`
	Phone         *string            `json:"Phone"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
