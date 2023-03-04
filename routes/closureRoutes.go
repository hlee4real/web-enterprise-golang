package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func ClosureRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/closures", controllers.CreateClosure())
	incomingRoutes.GET("/closures", controllers.GetAllClosures())
	incomingRoutes.GET("/closures/:id", controllers.GetClosureById())
	incomingRoutes.PUT("/closures/:id", controllers.UpdateClosure())
	incomingRoutes.DELETE("/closures/:id", controllers.DeleteClosure())
}
