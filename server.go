package main

import "github.com/gin-gonic/gin"

//logger
//routes
//controller
func main() {

	server := gin.Default()
	server.GET("/trdt", func(ctx *gin.Context)){
		ctx.JSON(200, gin.H{
			"mesage":"ok!"
		})
	}

}
