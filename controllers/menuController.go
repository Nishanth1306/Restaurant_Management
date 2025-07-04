package controller

import (
	"RestaurantManagement/database"
	"RestaurantManagement/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(ctx, bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing all the menus"})

		}
		var allmenus []bson.M
		if err = result.All(ctx, &allmenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allmenus)
	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured While fetching the Menu"})
			return
		}
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		now := time.Now()
		menu.Created_at = &now
		menu.Updated_at = &now

		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, inserterr := menuCollection.InsertOne(ctx, menu)
		if inserterr != nil {
			msg := fmt.Sprintf("Menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D
		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				msg := "kindly retype the time "
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				defer cancel()
				return
			}

			updateObj = append(updateObj, bson.E{"start_Date", menu.Start_Date})
			updateObj = append(updateObj, bson.E{"end_Date", menu.End_Date})

			if menu.Name != "" {
				updateObj = append(updateObj, bson.E{"name", menu.Name})
			}
			if menu.Category != "" {
				updateObj = append(updateObj, bson.E{"category", menu.Category})
			}
			
			now := time.Now()
			menu.Updated_at = &now
			updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result, err := menuCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"$set", updateObj},
				},
				&opt,
			)
			if err != nil {
				msg := fmt.Sprintf("Menu item was not updated")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})

			}
			defer cancel()
			c.JSON(http.StatusOK, result)
		}

	}
}
