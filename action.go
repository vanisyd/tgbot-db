package database

import (
	"context"
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
