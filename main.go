package main
import (
	"os"
	"github.com/gin-gonic/gin"
	"RestaurantManagement/database"
	"RestaurantManagement/routes"
	"RestaurantManagement/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	middleware "RestaurantManagement/middleware"
	
	
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"

	}
	router := gin.New()
	router.use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuItems(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":"+ port)


}