package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.GET("/users", controllers.GetAllUsers())
	incomingRoutes.GET("/users/:id", controllers.GetUserById())
	incomingRoutes.PUT("/users/:id", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/:id", controllers.DeleteUser())
}
