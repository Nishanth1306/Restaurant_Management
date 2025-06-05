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
)

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoicess() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(ctx, bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error Occured while fetching the invoices"})
		}

		var allinvoives []bson.M
		if err = result.All(ctx, &allinvoives); err != nil {
			log.Fatal(err)

		}
		c.JSON(http.StatusOK, allinvoives)
	}
}
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured While fetching the Invoice"})
			return
		}
		c.JSON(http.StatusOK, invoice)
	}
}
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var invoice models.Invoice
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		invoice.Invoice_id = primitive.NewObjectID().Hex()
		invoice.Created_at = time.Now()
		invoice.Updated_at = time.Now()
		invoice.ID = primitive.NewObjectID()
		validationErr := validate.Struct(invoice)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		result, insertErr := invoiceCollection.InsertOne(ctx, invoice)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while inserting the invoice"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "Invoice created successfully", "invoice_id": result.InsertedID})
		log.Println("Invoice created successfully:", result.InsertedID)
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
