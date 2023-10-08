package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	TgID int                `bson:"tg_id"`
}

func (u User) collectionName() string {
	return UsersCollection
}

func AddUser(user User) interface{} {
	coll := getCollection(user)
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted user with _id: %v\n", result.InsertedID)

	if uid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return uid
	}

	return nil
}

func FindUser(userID int) (result User) {
	coll := getCollection(User{})
	filter := bson.D{{"tg_id", userID}}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{}
	}

	return
}

func GetUser(userObjID primitive.ObjectID) (result User) {
	coll := getCollection(User{})
	filter := bson.D{{"_id", userObjID}}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return User{}
	}

	return
}
