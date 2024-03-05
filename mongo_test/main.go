package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	GUID string `bson:"guid"`
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("test")

	users := db.Collection("users")

	//user := User{GUID: "123"}
	//_, err = users.InsertOne(context.TODO(), user)
	//if err != nil {
	//	log.Fatal(err)
	//}

	filter := bson.D{{"guid", "123"}}
	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cursor.RemainingBatchLength())

	var result []User
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("all good")
}
