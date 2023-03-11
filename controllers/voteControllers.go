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

var userVoteCollection = models.GetUserVoteCollection()

func UpVote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var newVote models.UserVoteModels
		if err := c.BindJSON(&newVote); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		filter := bson.D{{"username", newVote.Username}, {"idea_id", newVote.IdeaId}}
		count, err := userVoteCollection.CountDocuments(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		if count == 0 {
			newVote.ID = primitive.NewObjectID()
			newVote.UpVote = true
			newVote.DownVote = false
			newVote.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			newVote.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			_, err = userVoteCollection.InsertOne(ctx, newVote)
			if err != nil {
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
		} else {
			update := bson.M{"$set": bson.M{"up_vote": true, "down_vote": false}}
			_, err = userVoteCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
		}
		c.Status(http.StatusOK)
	}
}

func DownVote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var vote models.UserVoteModels
		err := userVoteCollection.FindOne(ctx, bson.M{"username": vote.Username, "idea_id": vote.IdeaId}).Decode(&vote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		if vote.DownVote == true {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Already Voted", Data: nil})
			return
		}

		if vote.UpVote == true {
			//update vote
			_, err = userVoteCollection.UpdateOne(ctx, bson.M{"username": vote.Username, "idea_id": vote.IdeaId}, bson.M{"$set": bson.M{"up_vote": false, "down_vote": true}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: vote})
			return
		}

		var newVote models.UserVoteModels
		if err = c.BindJSON(&newVote); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		newVote.ID = primitive.NewObjectID()
		newVote.UpVote = false
		newVote.DownVote = true
		newVote.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newVote.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_, err = userVoteCollection.InsertOne(ctx, newVote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: newVote})
	}
}
