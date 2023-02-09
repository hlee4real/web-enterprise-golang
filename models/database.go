package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://hoanglh:13112002@backenddb.xjjfx9h.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		return nil, err
	}

	//set unique for username and index it in database
	mongoClient.Database("enterpriseweb").Collection("users").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{
			Key:   "username",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	})

	mongoClient.Database("enterpriseweb").Collection("departments").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{
			Key:   "name",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	})

	mongoClient.Database("enterpriseweb").Collection("documents").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{
			Key:   "filename",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	})

	mongoClient.Database("enterpriseweb").Collection("categories").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{
			Key:   "name",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	})

	mongoClient.Database("enterpriseweb").Collection("ideas").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{
			Key:   "title",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	})

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
