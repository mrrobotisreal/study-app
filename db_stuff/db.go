package db_stuff

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"study-app/auth"
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

func AddCollection(req AddCollectionRequest) bool {
	var user User
	var newCollection CardCollection
	err := users.FindOne(ctx, User{Username: req.Username}).Decode(&user)
	if err != nil {
		log.Println("Error Performing FindOne Operation | No Document Found..\n\n")
		return false
	}
	newCollection.CardList = req.CardList
	newCollection.Name = req.Name
	newCollection.Category = req.Category
	newCollection.Description = req.Description
	user.Collections = append(user.Collections, newCollection)
	filter := bson.D{{"username", req.Username}}
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
		log.Println(err)
	}
	return user.Collections
}

func GetCollection(username string, collectionName string) CardCollection {
	var user User
	err := users.FindOne(ctx, User{Username: username}).Decode(&user)
	if err != nil {
		log.Println("Errors be here.... oof....")
	}
	for i := 0; i < len(user.Collections); i++ {
		if user.Collections[i].Name == collectionName {
			return user.Collections[i]
		}
	}
	return CardCollection{}
}

func SetUser(user User) bool {
	log.Println("SetUser user is:")
	log.Println(user)
	_, err := users.InsertOne(ctx, user)
	if err != nil {
		log.Println("Error Performing InsertOne Operation..\n\n", err.Error())
		return false
	}
	return true
}

func GetUser(username string) User {
	var user User
	err := users.FindOne(ctx, User{Username: username}).Decode(&user)
	if err != nil {
		log.Println("Error Performing FindOne Operation | No Document Found..\n\n")
	}
	return user
}

func CheckLogin(username string, password string) LoginResult {
	var user User
	err := users.FindOne(ctx, User{Username: username}).Decode(&user)
	if err != nil {
		log.Println("Error Performing FindOne Operation | No Document Found..\ndb_stuff/db.go\nCheckLogin\n--------------------------------------------------------")
	}
	result := auth.CheckPasswordHash(password, user.Password)
	jwt := auth.GenerateJWT(user.Username, user.Email)
	var rslt LoginResult
	rslt.AuthResult = result
	rslt.JWT = jwt
	return rslt
}

func Logout(username string) bool {
	log.Println("Username inside db_stuff.Logout-> " + username)
	jwt := ""
	filter := bson.D{{"username", username}}
	update := bson.D{{"$set", bson.D{{"jwt", jwt}}}}
	rslt := users.FindOneAndUpdate(ctx, filter, update)
	log.Println("Rslt inside of db_stuff.Logout")
	log.Println(rslt)
	return true
}

func SetSession(username string, sessionReq SessionRequest) Session {
	var user User
	var jwt string
	jwt = auth.GenerateJWT(username, sessionReq.Email)
	log.Println("Generated JWT successfully")
	if sessionReq.IsSignup {
		log.Println("Inside if sessionReq.IsSignup")
		user.Username = username
		user.Email = sessionReq.Email
		user.Name = sessionReq.Name
		hashedPassword, err := auth.HashPassword("google-signup:" + sessionReq.Email)
		user.Password = hashedPassword
		if err != nil {
			log.Println("Error hashing password inside of SetSession")
			log.Println(err.Error())
		}
		log.Println("HashedPassword complete without error")
		user.JWT = jwt
		user.Collections = []CardCollection{}
		addUserResult := SetUser(user)
		log.Println("AddUserResult")
		log.Println(addUserResult)
		if !addUserResult {
			log.Println("FAIL")
		}
	} else {
		err := users.FindOne(ctx, User{Username: username}).Decode(&user)
		if err != nil {
			log.Println("Error Performing FindOne Operation | No Document Found..\ndb_stuff/db.go\nSetSession\n--------------------------------------------------------")
			log.Println(err.Error())
		}
		filter := bson.D{{"username", username}}
		update := bson.D{{"$set", bson.D{{"jwt", jwt}}}}
		rslt := users.FindOneAndUpdate(ctx, filter, update)
		log.Println("RESULT of update")
		log.Println(&rslt)
	}
	log.Println("JWT")
	log.Println(jwt)
	var session Session
	session.JWT = jwt
	return session
}
