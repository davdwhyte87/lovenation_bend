package main

import (
	"lovenation_bend/configs"
	"lovenation_bend/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	router:= gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data":"Hello bro we back here again"})
	})

	//run database
	configs.ConnectDB()

	// routes setup 
	routes.UserRoute(router)

	router.Run("localhost:6000")
}