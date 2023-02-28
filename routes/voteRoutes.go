package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func VoteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/up-vote", controllers.UpVote())
	incomingRoutes.POST("/down-vote", controllers.DownVote())
}
