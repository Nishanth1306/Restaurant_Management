package routes

import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"

)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
	incomingRoutes.POST("users/signup", controller.SignUp())
	incomingRoutes.PATCH("users/login", controller.Login())


}