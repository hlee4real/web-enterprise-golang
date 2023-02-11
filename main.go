package main

import (
	"github.com/gin-gonic/gin"
	middleware "web-enterprise-backend/middleware"
	routes "web-enterprise-backend/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

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
