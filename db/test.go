package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Add(p bson.D){
	
	collectionStudent := DB.Mongo.Database("go_web").Collection("affair")

	// doc := bson.D{p}
	
	_,err := collectionStudent.InsertOne(context.TODO(),p)
	if err != nil {
    	fmt.Println(err)
	}
}

func Find() any {
	collectionStudent := DB.Mongo.Database("go_web").Collection("affair")
	cursor ,err:= collectionStudent.Find(context.TODO(),bson.M{})
	if err != nil {
    	fmt.Println(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	 }
	for _, result := range results {
		fmt.Println(result)
	 }
	return results
} 