package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func StatisticRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ideas-submitted-by-departments/:department", controllers.IdeasSubmittedByDepartments())
	incomingRoutes.GET("/total-number-of-idea-submitted", controllers.TotalNumberOfIdeaSubmitted())
	incomingRoutes.GET("/total-number-of-idea-submitted-by-category/:category", controllers.TotalNumberOfIdeaSubmittedByCategory())
	incomingRoutes.GET("/highest-upvote-idea", controllers.GetHighestUpvoteIdea())
	incomingRoutes.GET("/highest-downvote-idea", controllers.GetHighestDownvoteIdea())
	incomingRoutes.GET("/staff-members-who-submitted-most-ideas", controllers.StaffMembersWhoSubmittedMostIdeas())
}
