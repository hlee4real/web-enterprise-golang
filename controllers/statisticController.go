package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"sort"
	"time"
	"web-enterprise-backend/models"
)

func IdeasSubmittedByDepartments() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		department := c.Param("department")
		//get all ideas of that department
		var ideas []models.IdeasModel
		cursor, err := ideaCollection.Find(ctx, bson.M{"department": department})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		if err = cursor.All(ctx, &ideas); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		numberOfIdeas := len(ideas)
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: numberOfIdeas})
	}
}

func TotalNumberOfIdeaSubmitted() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var ideas []models.IdeasModel
		cursor, err := ideaCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		if err = cursor.All(ctx, &ideas); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		numberOfIdeas := len(ideas)
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: numberOfIdeas})
	}
}

func TotalNumberOfIdeaSubmittedByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		category := c.Param("category")

		var ideas []models.IdeasModel
		cursor, err := ideaCollection.Find(ctx, bson.M{"category": category})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		if err = cursor.All(ctx, &ideas); err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		numberOfIdeas := len(ideas)
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: numberOfIdeas})
	}
}

func GetHighestUpvoteIdea() gin.HandlerFunc {
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
		sort.Slice(ideas, func(i, j int) bool {
			return ideas[i].UpVote > ideas[j].UpVote
		})
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: ideas[0]})
	}
}

func GetHighestDownvoteIdea() gin.HandlerFunc {
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
		sort.Slice(ideas, func(i, j int) bool {
			return ideas[i].DownVote > ideas[j].DownVote
		})
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: ideas[0]})
	}
}

func StaffMembersWhoSubmittedMostIdeas() gin.HandlerFunc {
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
		//count each username in the ideas array
		count := make(map[string]int)
		for _, idea := range ideas {
			count[idea.Username]++
		}

		//sort the count
		type kv struct {
			Key   string
			Value int
		}

		var ss []kv
		for k, v := range count {
			ss = append(ss, kv{k, v})
		}

		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value > ss[j].Value
		})
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: ss[0]})
	}
}
