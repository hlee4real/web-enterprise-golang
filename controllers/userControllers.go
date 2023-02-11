package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"web-enterprise-backend/helper"
	"web-enterprise-backend/models"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var userCollection = models.GetUserCollection()
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		return false, msg
	}
	return true, msg
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user models.UsersModel
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password
		count, err = userCollection.CountDocuments(ctx, bson.M{"mobile": user.Mobile})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mobile or username already exists"})
			return
		}
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Username, *user.FirstName, *user.LastName, user.UserID)
		user.Token = &token
		user.RefreshToken = &refreshToken
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": resultInsertionNumber})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.UsersModel
		var foundUser models.UsersModel
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Username, *foundUser.FirstName, *foundUser.LastName, foundUser.UserID)
		helper.UpdateAllTokens(token, refreshToken, foundUser.UserID)
		c.JSON(http.StatusOK, gin.H{"user": foundUser})
	}
}
