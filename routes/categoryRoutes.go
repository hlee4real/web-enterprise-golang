package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func CategoryRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/categories", controllers.GetAllCategories())
	incomingRoutes.GET("/categories/:id", controllers.GetCategoryById())
	incomingRoutes.POST("/categories", controllers.CreateCategory())
	incomingRoutes.PUT("/categories/:id", controllers.UpdateCategory())
	incomingRoutes.DELETE("/categories/:id", controllers.DeleteCategory())
}
