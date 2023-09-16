package main

import (
	// "context"
	"lovenation_bend/configs"
	// "lovenation_bend/dao"
	"lovenation_bend/routes"

	"github.com/gin-gonic/gin"
)

// var (
// 	FactoryDAO  *dao.FactoryDAO

// )
func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "Armorgens API running ..."})
	})

	//run database
	configs.ConnectDB()

	//
	// FactoryDAO = dao.InitializeFactory(configs.DB, context.TODO())

	// setup data access objects
	routes.SetupDAO()
	// routes setup
	routes.UserRoute(router)
	routes.VisaRoutes(router)

	router.Run("localhost:6000")
}
