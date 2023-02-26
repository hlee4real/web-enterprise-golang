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

var commentCollection = models.GetCommentCollection()

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var comment models.CommentsModel
		if err := c.BindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		comment.ID = primitive.NewObjectID()
		comment.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		comment.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_, err := commentCollection.InsertOne(ctx, comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: comment})
	}
}

func GetCommentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		var comment models.CommentsModel
		if err := commentCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&comment); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: comment})
	}
}

func GetAllComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var comments []models.CommentsModel
		cursor, err := commentCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		for cursor.Next(ctx) {
			var comment models.CommentsModel
			cursor.Decode(&comment)
			comments = append(comments, comment)
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: comments})
	}
}

func UpdateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		var comment models.CommentsModel
		if err := c.BindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		comment.ID = objectId
		comment.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		comment.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		if _, err := commentCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": comment}); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: comment})
	}
}

func DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		if _, err := commentCollection.DeleteOne(ctx, bson.M{"_id": objectId}); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: nil})
	}
}
