package main

import "github.com/gin-gonic/gin"

func main(){
	router:= gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data":"Hello bro we back here again"})
	})

}