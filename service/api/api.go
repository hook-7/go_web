package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"web/database"
)



func GetApi() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": fmt.Sprintf("https://%s:35212/", get_external()),
		})
	})

	router.GET("/ip", func(c *gin.Context)  {
	database.GetDB()
		c.JSONP(http.StatusOK, c.RemoteIP())
	})

	router.Run(":8080")
}

func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	return string(content)
}
