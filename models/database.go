package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://hoanglh:13112002@backenddb.xjjfx9h.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		return nil, err
	}

	return mongoClient, nil
}

func GetUserCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("users")
}

func GetCategoryCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		fmt.Println("error Database")
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("categories")
}

func GetDepartmentCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("departments")
}

func GetDocumentCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("documents")
}

func GetIdeaCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("ideas")
}

func GetCommentCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("comments")
}

func GetUserVoteCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("uservotes")
}

func GetClosureCollection() *mongo.Collection {
	mongoClient, err := ConnectDB()
	if err != nil {
		return nil
	}

	return mongoClient.Database("enterpriseweb").Collection("closures")
}
