package routes

import (
	"lovenation_bend/controllers"

	"github.com/gin-gonic/gin"
)

var visaController *controllers.VisaController

func initControllers(){
	factoryDAO := GetDAO()
	visaController = &controllers.VisaController{
		FactoryDAO: factoryDAO,
	}
}
func VisaRoutes(router *gin.Engine){
	initControllers()
	router.POST("/visa_application", visaController.CreateVisaApplication())	
}