package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"web/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetApi() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": fmt.Sprintf("https://%s:35212/", get_external()),
		})
	})

	router.GET("/ip", func(c *gin.Context) {
		coll := database.GetDB().Database("test").Collection("IOC_data")
		var result bson.M
		err := coll.FindOne(context.TODO(), bson.D{{Key: "test",Value:  "test"}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the title %s\n", "title")
			return
		}
		if err != nil {
			panic(err)
		}
		// coll.InsertOne(context.TODO(),bson.D{{Key: "test",Value: "test"},{Key: "test2",Value: "test2"}})
		jsonData, err := json.MarshalIndent(result, "", "    ")
		fmt.Printf("%s\n", jsonData)
		c.JSONP(http.StatusOK, result)
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
