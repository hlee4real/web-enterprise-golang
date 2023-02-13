package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func DocumentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/documents", controllers.GetAllDocuments())
	incomingRoutes.GET("/documents/:id", controllers.GetDocumentById())
	incomingRoutes.POST("/documents", controllers.CreateDocument())
	incomingRoutes.PUT("/documents/:id", controllers.UpdateDocument())
	incomingRoutes.DELETE("/documents/:id", controllers.DeleteDocument())
}
