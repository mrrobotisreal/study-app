package db_stuff

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var users *mongo.Collection
var ctx = context.TODO()

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	users = client.Database("flash-fire-webapp").Collection("users")
}

func AddCollection(username string, newCollection CardCollection) bool {
	var user User
	err := users.FindOne(ctx, User{Username: username}).Decode(&user)
	if err != nil {
		log.Println("Error Performing FindOne Operation | No Document Found..\n\n")
		return false
	}
	user.Collections = append(user.Collections, newCollection)
	filter := bson.D{{"username", username}}
	update := bson.D{{"$set", bson.D{{"collections", user.Collections}}}}
	result, err := users.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error Performing UpdateOne Operation..\n\n", err.Error())
		return false
	}
	log.Println(result)
	return true
}

func GetCollections(username string) []CardCollection {
	var user User
	err := users.FindOne(ctx, User{Username: username}).Decode(&user)
	if err != nil {
		log.Println("Errors be here...")
	}
	return user.Collections
}
