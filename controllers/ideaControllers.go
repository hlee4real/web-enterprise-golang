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

var ideaCollection = models.GetIdeaCollection()

func CreateIdea() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var idea models.IdeasModel
		if err := c.BindJSON(&idea); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		idea.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		idea.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_, err := ideaCollection.InsertOne(ctx, idea)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: idea})
	}
}

func GetIdeaById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		var idea models.IdeasModel
		if err := ideaCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&idea); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: idea})
	}
}

func GetAllIdeas() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var ideas []models.IdeasModel
		cursor, err := ideaCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		for cursor.Next(ctx) {
			var idea models.IdeasModel
			cursor.Decode(&idea)
			ideas = append(ideas, idea)
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: ideas})
	}
}

func UpdateIdea() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		var idea models.IdeasModel
		if err := c.BindJSON(&idea); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		idea.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		idea.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		if _, err := ideaCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": idea}); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: idea})
	}
}

func DeleteIdea() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		if _, err := ideaCollection.DeleteOne(ctx, bson.M{"_id": objectId}); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: "Deleted idea"})
	}
}
