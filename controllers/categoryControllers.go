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

var categoryCollection = models.GetCategoryCollection()

func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var category models.CategoriesModel
		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		category.ID = primitive.NewObjectID()
		category.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		category.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_, err := categoryCollection.InsertOne(ctx, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: category})
	}
}

func GetCategoryById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		var category models.CategoriesModel
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}

		err = categoryCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: category})
	}
}

func GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var categories []models.CategoriesModel
		cursor, err := categoryCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		for cursor.Next(ctx) {
			var category models.CategoriesModel
			cursor.Decode(&category)
			categories = append(categories, category)
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: categories})
	}
}

func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		var category models.CategoriesModel
		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		category.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		category.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		_, err := categoryCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": category})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: category})
	}
}

func DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		_, err := categoryCollection.DeleteOne(ctx, bson.M{"_id": objectId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIResponse{Status: 0, Message: "Error", Data: nil})
			return
		}
		c.JSON(http.StatusOK, APIResponse{Status: 1, Message: "Success", Data: "Delete Success"})
	}
}
