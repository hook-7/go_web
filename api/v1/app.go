package v1

import (
	"github.com/gin-gonic/gin"
	"go_web/db"
	"fmt"
)

func Test(ctx *gin.Context) {
	
	p := mongodb.Find()
	fmt.Println(&p)
	ctx.JSON(200, &p)
}