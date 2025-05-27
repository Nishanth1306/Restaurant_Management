// package helpers

// import (
// 	"RestaurantManagement/database"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	jwt "github.com/golang-jwt/jwt/v4"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type SignedDetails struct {
// 	Email      string
// 	First_name string
// 	Last_name  string
// 	Uid        string
// 	jwt.StandardClaims
// }

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
// var Secret_Key string = os.Getenv("SECRET_KEY")

// func GenerateAllTokens(email, firstName, lastName, uid string) (signedToken string, signedRefreshToken string, err error) {
// 	claims := &SignedDetails{
// 		Email:      email,
// 		First_name: firstName,
// 		Last_name:  lastName,
// 		Uid:        uid,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
// 			Issuer:    "RestaurantManagement",
// 		},
// 	}

// 	refreshClains := &SignedDetails{
// 		Email:      email,
// 		First_name: firstName,
// 		Last_name:  lastName,
// 		Uid:        uid,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
// 			Issuer:    "RestaurantManagement",
// 		},
// 	}
// 	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Secret_Key))
// 	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClains).SignedString([]byte(Secret_Key))

// 	if err != nil {
// 		log.Panic("Error while generating token")
// 	}

// 	return token, refreshToken, err
// }

// func UpdateAllTokens() {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	var updateObj primitive.D

// 	updateObj = append(updateObj, bson.E{"token", signedToken})
// 	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

// 	Update_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

// 	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Update_at})

// 	upsert := true
// 	filter := bson.M{"user_id": userId}
// 	opt := options.UpdateOptions{
// 		Upsert: &upsert,
// 	}

// 	result, err := userCollection.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.D{
// 			{"$set", updateObj},
// 		},
// 		&opt,
// 	)

// 	defer cancel()

// 	if err != nil {
// 		log.Panic(err)
// 		return

// 	}

// 	return
// }

// func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {

// 	jwt.ParseWithClain(
// 		signedToken,
// 		&SignedDetails{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(Secret_Key), nil
// 		},
// 	)

// 	claims, ok := token.Claims.(*SignedDetails)
// 	if !ok {
// 		msg = fmt.Sprintf("the token is invalid")
// 		msg = err.Error()
// 	}

// 	if claims.ExpiresAt < time.Now().Local().Unix() {
// 		msg = fmt.Sprintf("the token is expired")
// 		msg = err.Error()
// 	}

// 	return claims, msg
// }

package helpers

import (
	"RestaurantManagement/database"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var Secret_Key string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email, firstName, lastName, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
			Issuer:    "RestaurantManagement",
		},
	}

	refreshClaims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
			Issuer:    "RestaurantManagement",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Secret_Key))
	if err != nil {
		log.Panic("Error while generating token")
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(Secret_Key))
	if err != nil {
		log.Panic("Error while generating refresh token")
		return
	}

	return token, refreshToken, err
}

// You need to pass required parameters to this function (e.g., signedToken, signedRefreshToken, userId)
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)

	if err != nil {
		log.Panic(err)
		return
	}

	return
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(Secret_Key), nil
		},
	)

	if err != nil {
		msg = fmt.Sprintf("the token is invalid: %v", err)
		return nil, msg
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok || !token.Valid {
		msg = fmt.Sprintf("the token is invalid")
		return nil, msg
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("the token is expired")
		return nil, msg
	}

	return claims, ""
}
