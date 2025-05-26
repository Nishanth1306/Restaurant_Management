package routes 
import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"
)
func OrderItemRoutes(incominRoutes *gin.Engine){
	incominRoutes.GET("/orderItems", controller.GetOrderItems())
	incominRoutes.GET("/orderItems/:orderItem_id",controller.GetOrderItem())
	incominRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incominRoutes.POST("/orderItems", controller.CreateOrderItem())
	incominRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}