package routes

import (
	"lovenation_bend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    router.POST("/user", controllers.CreateUser())
}
