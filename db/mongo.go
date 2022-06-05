package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
const uri = "mongodb://admin:123456@172.24.0.1:27017"


type Database struct {
	Mongo  *mongo.Client
}


var DB *Database


//初始化
func Init() {
	DB = &Database{
		Mongo: SetConnect(),
	}
}
// 连接设置
func SetConnect() *mongo.Client{
	ctx ,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	client, err :=      mongo.Connect(ctx,options.Client().ApplyURI(uri).SetMaxPoolSize(20))
	if err !=nil{
		fmt.Println(err)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client
}


