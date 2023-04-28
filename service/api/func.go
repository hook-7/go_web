package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net"
	"net/http"
	"strings"
	"time"
	"web/database"
)

type DDNS struct {
	Domain string `bson:"domain"`
	Date   string `bson:"date"`
}

var collection = database.GetDB().Database("test").Collection("DDNS")

func login(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	fmt.Println(param)
	fmt.Println(param["pwd"])
	if err != nil {
		return
	}
}

func page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "极简生活",
	})
}

func setip(c *gin.Context) {
	ip := strings.Split(c.Request.Header.Get("X-Forwarded-For"), ",")[0]
	if ip == "" {
		ip = c.Request.RemoteAddr
		ip, _, _ = net.SplitHostPort(ip)
	}
	ip = strings.TrimSpace(ip)

	filter := bson.M{"domain": ip}
	var result DDNS
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err == mongo.ErrNoDocuments {
		currentTime := time.Now().UTC()
		record := DDNS{Domain: ip, Date: currentTime.Format("2006-01-02 15:04:05")}
		_, err := collection.InsertOne(context.Background(), record)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.String(200, ip)
}
