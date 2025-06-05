package controller

import (
	"RestaurantManagement/database"
	"RestaurantManagement/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := tableCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing all the tables"})
			return
		}
		var allTables []bson.M
		if err = result.All(ctx, &allTables); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding the tables"})
			return
		}
		c.JSON(http.StatusOK, allTables)

	}
}
func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")
		var table models.Table
		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured While fetching the Table"})
			return
		}
		c.JSON(http.StatusOK, table)

	}
}
func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()

		result, inserterr := tableCollection.InsertOne(ctx, table)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while inserting the table"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Table created successfully", "table_id": result.InsertedID})

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		tableId := c.Param("table_id")
		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj primitive.D
		if table.Number_of_guests != nil {
			updateObj = append(updateObj, bson.E{"number_of_guests", table.Number_of_guests})
		}
		if table.Table_number != nil {
			updateObj = append(updateObj, bson.E{"table_number", table.Table_number})
		}
		updateObj = append(updateObj, bson.E{"updated_at", time.Now()})
		result, err := tableCollection.UpdateOne(ctx, bson.M{"table_id": tableId}, bson.D{{"$set", updateObj}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the table"})
			return
		}
		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Table updated successfully", "matched_count": result.MatchedCount})

	}
}
