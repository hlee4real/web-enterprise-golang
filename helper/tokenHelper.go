package helper

import (
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"web-enterprise-backend/models"
)

type SignedDetails struct {
	Username  string
	FirstName string
	LastName  string
	Uid       string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = models.GetUserCollection()
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(username string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		fmt.Println(err)
		return
	}
	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = "Invalid token"
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "Invalid token"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Token expired"
		return
	}
	return claims, msg
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", updatedAt})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{Upsert: &upsert}
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	defer cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
