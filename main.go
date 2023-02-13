package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"web-enterprise-backend/middleware"
	"web-enterprise-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	router := gin.Default()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.CategoryRoutes(router)
	routes.DocumentRoutes(router)
	routes.DepartmentRoutes(router)
	routes.IdeaRoutes(router)
	routes.CommentRoutes(router)
	router.Use(middleware.Authentication())

	router.Run(":8080")
}
