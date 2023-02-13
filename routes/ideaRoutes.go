package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func IdeaRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ideas", controllers.GetAllIdeas())
	incomingRoutes.GET("/ideas/:id", controllers.GetIdeaById())
	incomingRoutes.POST("/ideas", controllers.CreateIdea())
	incomingRoutes.PUT("/ideas/:id", controllers.UpdateIdea())
	incomingRoutes.DELETE("/ideas/:id", controllers.DeleteIdea())
}
