package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Bot struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Token  string             `bson:"token"`
	HashID string             `bson:"hash_id"`
}

func (b Bot) collectionName() string {
	return BotsCollection
}

func AddBot(bot Bot) interface{} {
	coll := getCollection(bot)
	result, err := coll.InsertOne(context.TODO(), bot)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted bot with _id: %v\n", result.InsertedID)

	if uid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return uid
	}

	return nil
}

func FindBot(hashID string) interface{} {
	var result Bot
	coll := getCollection(result)
	filter := bson.D{{"hash_id", hashID}}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil
	}

	return result
}

func GetBots() interface{} {
	var result []Bot
	coll := getCollection(Bot{})

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil
	}

	for cursor.Next(context.TODO()) {
		var botData Bot
		err := cursor.Decode(&botData)
		if err != nil {
			log.Println("[DB.Error]" + err.Error())
			continue
		}
		result = append(result, botData)
	}

	if err := cursor.Err(); err != nil {
		log.Println("[DB.Error]" + err.Error())
	}

	err = cursor.Close(context.TODO())
	if err != nil {
		log.Println("[DB.Error]" + err.Error())
	}
	return result
}
