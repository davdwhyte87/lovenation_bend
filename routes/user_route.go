package routes

import (
	"lovenation_bend/controllers"

	"github.com/gin-gonic/gin"
)


var (
	userController *controllers.UserController
)
func initializeController(){
	factoryDAO, _ := GetDAO()

	userController = &controllers.UserController{
		FactoryDAO: factoryDAO,
	
	}
}
func UserRoute(router *gin.Engine)  {
	// initialize controllers
	initializeController()

    router.POST("/user", userController.CreateUser())
}
