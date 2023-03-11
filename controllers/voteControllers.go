package controllers

import (
	"context"
	"fmt"
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
			fmt.Println(err)
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
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//get idea
			var idea models.IdeasModel
			ideaId, _ := primitive.ObjectIDFromHex(newVote.IdeaId)
			err := ideaCollection.FindOne(ctx, bson.M{"_id": ideaId}).Decode(&idea)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//update idea
			idea.ID = ideaId
			idea.UpVote += 1
			idea.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			_, err = ideaCollection.UpdateOne(ctx, bson.M{"_id": ideaId}, bson.M{"$set": idea})
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
		} else {
			//get user vote and check if is upvoting or downvoting
			var userVote models.UserVoteModels
			err := userVoteCollection.FindOne(ctx, filter).Decode(&userVote)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			if userVote.UpVote == true {
				c.JSON(http.StatusOK, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			update := bson.M{"$set": bson.M{"up_vote": true, "down_vote": false}}
			_, err = userVoteCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//get idea
			var idea models.IdeasModel
			ideaId, _ := primitive.ObjectIDFromHex(newVote.IdeaId)
			err = ideaCollection.FindOne(ctx, bson.M{"_id": ideaId}).Decode(&idea)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//update idea
			idea.ID = ideaId
			idea.UpVote += 1
			idea.DownVote -= 1
			idea.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			_, err = ideaCollection.UpdateOne(ctx, bson.M{"_id": ideaId}, bson.M{"$set": idea})
			if err != nil {
				fmt.Println(err)
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
		var newVote models.UserVoteModels

		if err := c.BindJSON(&newVote); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		filter := bson.D{{"username", newVote.Username}, {"idea_id", newVote.IdeaId}}
		count, err := userVoteCollection.CountDocuments(ctx, filter)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		if count == 0 {
			newVote.ID = primitive.NewObjectID()
			newVote.UpVote = false
			newVote.DownVote = true
			newVote.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			newVote.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			_, err = userVoteCollection.InsertOne(ctx, newVote)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//get idea
			var idea models.IdeasModel
			ideaId, _ := primitive.ObjectIDFromHex(newVote.IdeaId)
			err := ideaCollection.FindOne(ctx, bson.M{"_id": ideaId}).Decode(&idea)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//update idea
			idea.ID = ideaId
			idea.DownVote += 1
			idea.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			_, err = ideaCollection.UpdateOne(ctx, bson.M{"_id": ideaId}, bson.M{"$set": idea})
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
		} else {
			//get user vote and check if is upvoting or downvoting
			var userVote models.UserVoteModels
			err := userVoteCollection.FindOne(ctx, filter).Decode(&userVote)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			if userVote.DownVote == true {
				c.JSON(http.StatusOK, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			update := bson.M{"$set": bson.M{"up_vote": false, "down_vote": true}}
			_, err = userVoteCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//get idea
			var idea models.IdeasModel
			ideaId, _ := primitive.ObjectIDFromHex(newVote.IdeaId)
			err = ideaCollection.FindOne(ctx, bson.M{"_id": ideaId}).Decode(&idea)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
			//update idea
			idea.ID = ideaId
			idea.UpVote -= 1
			idea.DownVote += 1
			idea.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			_, err = ideaCollection.UpdateOne(ctx, bson.M{"_id": ideaId}, bson.M{"$set": idea})
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
				return
			}
		}
		c.Status(http.StatusOK)
	}
}
