package controller

import (
	"RestaurantManagement/database"
	"RestaurantManagement/models"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollection.Find(ctx, bson.M{})

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing all the order"})
		}
		var allorders []bson.M
		if err = result.All(ctx, &allorders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allorders)

	}
}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order bson.M
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured While fetching the Order"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)

		defer cancel()

		var order models.Order
		if err := c.BindJSON(&order);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return	
		}
		order.Order_id = primitive.NewObjectID().Hex()
		order.Created_at = time.Now()
		order.Updated_at = time.Now()
		order.ID = primitive.NewObjectID()

		order.Table_id = nil


		result, insertErr := orderCollection.InsertOne(ctx, order)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while inserting the order"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"order_id": result.InsertedID, "message": "Order created successfully"})

	}
}

// func UpdateOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var table models.Table
// 		var order models.Order

// 		var updateObj primitive.D
// 		orderId := c.Param("order_id")

// 		if err := c.BindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if order.Table_id != nil {
// 			err = menuCollection.FindOne(ctx, bson.M{"table_id": food.Table_id}).Decode(&table)
// 			defer cancel()
// 			if err != nil {
// 				msg := "Table was not found"
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 			updateObj = append(updateObj, bson.E{"table_id", table.Table_id})
// 		}
// 		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		updateObj = append(updateObj, bson.E{"updated_at", food.Updated_at})

// 	    upsert:= true

// 		filter := bson.M{"order_id": orderId}

// 		opt := options.UpdateOptions{
// 			Upsert: &upsert,
// 		}

// 		orderCollection.UpdateOne(
// 			ctx,
// 			filter,
// 			bson.D{
// 				{"$set", updateObj},
// 			},
// 			&opt,
// 		)
// 		if err != nil {
// 			msg := fmt.Sprintf("Order item was not updated")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		defer cancel()
// 		c.JSON(http.StatusOK, result)
// 	}
// }

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var table models.Table
		var order models.Order
		var updateObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Table was not found"})
				return
			}
			updateObj = append(updateObj, bson.E{"table_id", table.Table_id})
		}
		order.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

		upsert := true
		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{Upsert: &upsert}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{"$set", updateObj}},
			&opt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Order item was not updated"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
