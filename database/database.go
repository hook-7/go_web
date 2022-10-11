package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Singleton struct {
	db *mongo.Client
}


var once sync.Once
var db  *Singleton

func useDatabase () *mongo.Client{
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://by9559:2FjNWHfZNnoHcoIY@serverdb.tgnra.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("!!!!!~~~~~~~~~")
	return client
}



func GetDB() *Singleton {
	once.Do(func() {
		db = &Singleton{
			db:useDatabase(),
		}
	})
	return db
}