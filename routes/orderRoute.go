package routes 
import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"
)
func OrderRoutes(incominRoutes *gin.Engine){
	incominRoutes.GET("/orders", controller.GetOrders())
	incominRoutes.GET("/orders/:orders_id",controller.GetOrder())
	incominRoutes.POST("/orders", controller.CreateOrder())
	incominRoutes.PATCH("./orders/order_id", controller.UpdateOrder())
}