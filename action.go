package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Action struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId primitive.ObjectID `bson:"user_id"`
	Data   interface{}        `bson:"data"`
}

func (a Action) collectionName() string {
	return ActionsCollection
}

func AddAction(action Action) {
	coll := getCollection(action)
	result, err := coll.InsertOne(context.TODO(), action)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted action with _id: %v\n", result.InsertedID)
}

func FindActions(userID primitive.ObjectID) (result []Action) {
	coll := getCollection(Action{})
	filter := bson.D{{"user_id", userID}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil
	}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Fatal(err)
	}

	return
}
