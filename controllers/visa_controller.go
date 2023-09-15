package controllers

import (
	"context"
	"log"
	"lovenation_bend/configs"
	"lovenation_bend/dao"
	"lovenation_bend/models"
	"lovenation_bend/requests"
	"lovenation_bend/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var visaApplicationAnswersCollection *mongo.Collection = configs.GetCollection(configs.DB, "VisaApplicationAnswers")
var visaApplicationCollection *mongo.Collection = configs.GetCollection(configs.DB, "VisaApplications")


type VisaController struct{
	FactoryDAO *dao.FactoryDAO
}

func (visaController VisaController) CreateVisaApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var visaApplicationRequest requests.VisaApplicationRequest
		err := c.BindJSON(&visaApplicationRequest)

		// validate request body
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// validate data
		validationError := validate.Struct(&visaApplicationRequest)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation Error",
				Data:    map[string]interface{}{"data": validationError.Error()},
			})
			return
		}
		// create application id
		applicationId := primitive.NewObjectID()

		// assign id to all questions answers
		applicationAnswers := make([]interface{}, 0)
		for _, ans := range visaApplicationRequest.ApplicationAnswers {
			var applicationAnswer models.ApplicationAnswers

			applicationAnswer.ApplicationId = applicationId.String()
			applicationAnswer.Id = primitive.NewObjectID()
			applicationAnswer.QuestionId = ans.QuestionId
			applicationAnswer.TextAnswer = ans.TextAnswer
			applicationAnswer.YesNoAnswer = ans.YesNoAnswer
			applicationAnswers = append(applicationAnswers, applicationAnswer)

		}

		// create database context
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		log.Printf("ans", len(visaApplicationRequest.ApplicationAnswers))
		// insert all the answers to the questions
		// visaApplicationAnswersCollection.InsertMany(ctx, applicationAnswers)
		insertAnsErr := visaController.FactoryDAO.InsertMany(models.VisaApplicationAnswers, applicationAnswers)
		if insertAnsErr != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status: http.StatusInternalServerError,
				Message: "Error",
				Data: map[string]interface{}{"data":insertAnsErr.Error()},
			})
			return
		}

		// insert new application
		var visaApplication models.VisaApplication
		visaApplication.Id = applicationId
		visaApplication.Email = visaApplicationRequest.Email
		visaApplication.Phone = visaApplicationRequest.Phone
		visaApplication.Name = visaApplicationRequest.Name
		visaApplication.Location = visaApplicationRequest.Location
		visaApplication.Profession = visaApplicationRequest.Profession

		// result, _ := visaApplicationCollection.InsertOne(ctx, visaApplication)

		insertErr := visaController.FactoryDAO.Insert(models.VisaApplications, visaApplication )
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status: http.StatusInternalServerError,
				Message: "Error",
				Data: map[string]interface{}{"data":insertErr.Error()},
			})
			return
		}
		// send response
		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Created Application",
			Data:    map[string]interface{}{"data": "OK"},
		})
	}
}

func (visaController VisaController) ApproveVisaApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get id from request
		applicationId := c.Param("id")
		var visaApplication models.VisaApplication
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := visaApplicationCollection.FindOne(ctx, bson.M{"_id": applicationId}).Decode(&visaApplication)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//ctxUpdate, cancelUpdate := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancelUpdate()
		// update the visa application
		//updateResult, updateErr := visaApplicationCollection.UpdateOne(ctxUpdate, bson.M{"_id"}, &visaApplication)

	}
}
