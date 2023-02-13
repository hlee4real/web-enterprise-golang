package routes

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/controllers"
)

func DepartmentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/departments", controllers.GetAllDepartments())
	incomingRoutes.GET("/departments/:id", controllers.GetDepartmentById())
	incomingRoutes.POST("/departments", controllers.CreateDepartment())
	incomingRoutes.PUT("/departments/:id", controllers.UpdateDepartment())
	incomingRoutes.DELETE("/departments/:id", controllers.DeleteDepartment())
}
