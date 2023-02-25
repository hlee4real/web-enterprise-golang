package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"sort"
	"time"
	"web-enterprise-backend/helper"
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
		//send email to manager
		email := new(helper.Email)
		emailAddress := helper.GetManagerEmail(idea.Department)
		if emailAddress != "" {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		email.Receiver = emailAddress
		email.Subject = fmt.Sprintf("New idea: %s", idea.Title)
		email.Body = fmt.Sprintf("%s", idea.Content)
		err = helper.SendEmail(*email)
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
		pageFilter := new(helper.PageFilter)
		err := pageFilter.Check(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		findOptions := options.Find()
		findOptions.SetSkip((int64(pageFilter.Page) - 1) * int64(pageFilter.Size))
		findOptions.SetLimit(int64(pageFilter.Size))
		cursor, err := ideaCollection.Find(ctx, bson.M{}, findOptions)
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

func GetMostViewedIdeas() gin.HandlerFunc {
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
		//sort idea by views
		sort.Slice(ideas, func(i, j int) bool {
			return ideas[i].Views > ideas[j].Views
		})
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: ideas})
	}
}

func GetMostUpvoteIdeas() gin.HandlerFunc {
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
		//sort idea by upvote
		sort.Slice(ideas, func(i, j int) bool {
			return ideas[i].UpVote > ideas[j].UpVote
		})
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: ideas})
	}
}
