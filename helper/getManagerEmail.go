package helper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"web-enterprise-backend/models"
)

func GetManagerEmail(department string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var manager models.UsersModel
	//find with role manager and department
	err := userCollection.FindOne(ctx, bson.M{"role": "Manager", "department_id": department}).Decode(&manager)
	if err != nil {
		return ""
	}
	return *manager.Username
}
