package routes

import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controller.GetUsers())
	router.GET("/users/:user_id", controller.GetUser())
	router.POST("/users/signup", controller.SignUp())
	router.POST("/users/login", controller.Login())
}
