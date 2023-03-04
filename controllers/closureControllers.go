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

var closureCollection = models.GetClosureCollection()

func CreateClosure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var closure models.ClosureModels
		if err := c.BindJSON(&closure); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		closure.ID = primitive.NewObjectID()
		_, err := closureCollection.InsertOne(ctx, closure)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: closure})
	}
}

func GetClosureById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var closure models.ClosureModels
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		if err := closureCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&closure); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: closure})
	}
}

func GetAllClosures() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var closures []models.ClosureModels
		cursor, err := closureCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		if err = cursor.All(ctx, &closures); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: closures})
	}
}

func UpdateClosure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var closure models.ClosureModels
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		if err := c.BindJSON(&closure); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		closure.ID = objectId
		if err := closureCollection.FindOneAndReplace(ctx, bson.M{"_id": objectId}, closure).Decode(&closure); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: closure})
	}
}

func DeleteClosure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		if _, err := closureCollection.DeleteOne(ctx, bson.M{"_id": objectId}); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: nil})
	}
}
