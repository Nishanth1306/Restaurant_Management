package routes 
import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"
)
func TableRoutes(incominRoutes *gin.Engine){
	incominRoutes.GET("/tables", controller.GetTables())
	incominRoutes.GET("/tables/:table_id",controller.GetTable())
	incominRoutes.POST("/tables", controller.CreateTable())
	incominRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
}