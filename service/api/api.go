package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"web/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginForm struct {
	user     string `form:"user" binding:"required"`
	pwd string `form:"password" binding:"required"`
}

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
		err := coll.FindOne(context.TODO(), bson.D{{Key: "test", Value: "test"}}).Decode(&result)
		cursor, err := coll.Find(context.TODO(), bson.D{{Key: "test", Value: "test"}})
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the title %s\n", "title")
			return
		}
		if err != nil {
			panic(err)
		}
		// coll.InsertOne(context.TODO(),bson.D{{Key: "test",Value: "test"},{Key: "test2",Value: "test2"}})

		c.JSONP(http.StatusOK, results)
	})

	router.POST("/login", func(c *gin.Context) {
		param := make(map[string]interface{})
		err := c.BindJSON(&param)
			fmt.Println(param)
			fmt.Println(param["pwd"])
			if err != nil {
				return 
			}

	
	})

	router.Run(":8081")
}

func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	updateDns()
	return string(content)
}

func updateDns() {
	url := "https://api.cloudflare.com/client/v4/zones/bf2b8a3249fd02ec38d9f9eb17fe1e37/dns_records/c30054c61bf1a6d9db6d053667b8f821"
    payload := make(map[string]interface {})
	payload["content"] = "66.150.128.82"
	payload["name"] = "test"
	payload["proxied"] = true
	payload["proxiable"] = true
	payload["ttl"] = 1
	payload["type"] = "A"
	marshal, _ := json.Marshal(payload)
	reader := bytes.NewReader(marshal)
	req, _ := http.NewRequest("PUT", url, reader)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Email", "war7ng@qq.com")
	req.Header.Add("X-Auth-Key","45aac82ead6b0e6df93fed2bcfaf54ad74412")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	
}


