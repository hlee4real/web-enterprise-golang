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
	router.Use(middleware.Authentication())

	router.GET("/api-v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "access grant for api-v1",
		})
	})

	router.GET("/api-v2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "access grant for api-v2",
		})
	})

	router.Run(":8080")
}
