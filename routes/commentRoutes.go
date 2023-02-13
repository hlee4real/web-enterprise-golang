package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func CommentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/comments", controllers.GetAllComments())
	incomingRoutes.GET("/comments/:id", controllers.GetCommentById())
	incomingRoutes.POST("/comments", controllers.CreateComment())
	incomingRoutes.PUT("/comments/:id", controllers.UpdateComment())
	incomingRoutes.DELETE("/comments/:id", controllers.DeleteComment())
}
