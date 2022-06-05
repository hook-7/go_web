package main

import (
	"fmt"
	"go_web/api/v1"
	"go_web/db"

	"github.com/gin-gonic/gin"
)


func main() {
	
	// mongodb.InitDb()
	mongodb.Init()

	router := gin.Default()

	router.GET("/:name/:age/",func (c *gin.Context){
		p :=c.Params
		c.JSON(200, p)
		fmt.Println(p)
	
		// mongodb.Add(p)
	})

	router.POST("/test", v1.Test)
	router.Run(":8080")
	
}