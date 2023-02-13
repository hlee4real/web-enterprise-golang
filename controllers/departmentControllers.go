package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
	"web-enterprise-backend/models"
)

var departmentCollection = models.GetDepartmentCollection()

func CreateDepartment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var department models.DepartmentsModel
		if err := c.BindJSON(&department); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		department.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		department.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_, err := departmentCollection.InsertOne(ctx, department)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: department})
	}
}

func GetDepartmentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		var department models.DepartmentsModel
		objectId, _ := primitive.ObjectIDFromHex(id)
		if err := departmentCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&department); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: department})
	}
}

func GetAllDepartments() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var departments []models.DepartmentsModel
		cursor, err := departmentCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		for cursor.Next(ctx) {
			var department models.DepartmentsModel
			cursor.Decode(&department)
			departments = append(departments, department)
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: departments})
	}
}

func UpdateDepartment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		var department models.DepartmentsModel
		objectId, _ := primitive.ObjectIDFromHex(id)
		if err := c.BindJSON(&department); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		department.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		department.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		if _, err := departmentCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": department}); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: department})
	}
}

func DeleteDepartment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		if _, err := departmentCollection.DeleteOne(ctx, bson.M{"_id": objectId}); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: "deleted"})
	}
}
