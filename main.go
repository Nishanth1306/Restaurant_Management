package main

import (
	"RestaurantManagement/database"
	"RestaurantManagement/metrics"
	"RestaurantManagement/middleware"
	"RestaurantManagement/routes"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func metricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		metrics.HttpRequestsTotal.WithLabelValues(c.FullPath(), c.Request.Method).Inc()
		metrics.HttpRequestDuration.WithLabelValues(c.FullPath()).Observe(duration.Seconds())
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()

	metrics.Register()

	router.Use(metricsMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
