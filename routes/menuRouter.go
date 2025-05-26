package routes
  
import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"
)

func MenuRoutes(incominRoutes *gin.Engine){

	incominRoutes.GET("/menus", controller.GetMenus())
	incominRoutes.GET("/menus/:menu_id",controller.GetMenu())
	incominRoutes.POST("/menus", controller.CreateMenu())
	incominRoutes.PATCH("./menuss/menu_id", controller.UpdateMenu())



}