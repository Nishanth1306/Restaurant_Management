package routes 
import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"
)
func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:orders_id",controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("./orders/order_id", controller.UpdateOrder())
}