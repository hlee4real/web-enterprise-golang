package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-enterprise-backend/models"
)

type RegisterInput struct {
	Username     string `bson:"username" json:"username" binding:"required"`
	Password     string `bson:"username" json:"password" binding:"required"`
	DateOfBirth  string `bson:"date_of_birth" json:"date_of_birth"`
	Mobile       string `bson:"mobile" json:"mobile"`
	Role         string `bson:"role" json:"role"`
	Image        string `bson:"image" json:"image"`
	DepartmentId string `bson:"department_id" json:"department_id"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.UsersModel{}
	u.Username = input.Username
	u.Password = input.Password
	u.DateOfBirth = input.DateOfBirth
	u.Mobile = input.Mobile
	u.Role = input.Role
	u.Image = input.Image
	u.DepartmentId = input.DepartmentId
	u.SaveUser()
	c.JSON(http.StatusOK, gin.H{"message": "validated!"})
}
