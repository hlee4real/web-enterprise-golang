package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
}
