package main

import (
	"github.com/gin-gonic/gin"
	"web-enterprise-backend/middleware"
)

func main() {
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", middleware.Register)
	r.Run(":8080")
}
