package controllers

import (
	// "context"
	"lovenation_bend/configs"
	"lovenation_bend/dao"
	"lovenation_bend/models"
	"lovenation_bend/responses"
	// "time"

	"net/http"

	// "firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "Users")


type UserController struct{
	FactoryDAO *dao.FactoryDAO	
}

func (userController UserController) CreateUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)

		
		// validate request body 
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status: http.StatusBadRequest,
				Message: "Error",
				Data: map[string]interface{}{"data": err.Error()},
			})
		}

		// validate data 
		validationError := validate.Struct(&user)
		if validationError != nil{
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status: http.StatusBadRequest,
				Message: "Validation Error",
				Data: map[string]interface{}{"data":validationError.Error()},
			})
		}

		// create timeout context for db
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()

		// prep user data 
		newUser := models.User{
			Id: primitive.NewObjectID(),
			FullName: user.FullName,
			Location: user.Location,
			
		}

		// insert user data into db
		// result, err := userCollection.InsertOne(ctx,newUser )
		insertErr := userController.FactoryDAO.Insert("Users", newUser)
		if insertErr != nil {
            c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

		
		// return data 
		c.JSON(http.StatusOK, responses.UserResponse{
			Status: http.StatusOK,
			Data: map[string]interface{}{"data":"OK"},
			Message: "User created",
		})

	}
}