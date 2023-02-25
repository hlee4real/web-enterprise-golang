package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/helper"
)

func UploadFileRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/upload-file", helper.UploadFile)
}
