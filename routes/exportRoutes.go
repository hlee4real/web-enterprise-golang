package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func ExportRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/export-ideas-to-csv", controllers.ExportIntoCSV())
	incomingRoutes.GET("/export-documents-to-zip", controllers.ExportIntoZip())
}
