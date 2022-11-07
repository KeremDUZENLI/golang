package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/mongoDatabase"))
	if err != nil {
		log.Fatal(err)
	}

	timeOut, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(timeOut)
	defer client.Disconnect(timeOut)

	mongoDatabase := client.Database("mongoDatabase")
	mongoCollection := mongoDatabase.Collection("mongoCollection")
	mongoEpisodes := mongoDatabase.Collection("mongoEpisodes")
	// mongoDatabase.CreateCollection(timeOut, "teste", nil)

	// mongoCollection - InsertOne
	valuesJson_1 := bson.D{
		{Key: "key_1", Value: "value_1"},
		{Key: "key_2", Value: "value_2"},
	}

	data1, _ := mongoCollection.InsertOne(timeOut, valuesJson_1)

	fmt.Println(data1.InsertedID)

	// mongoEpisodes - InsertMany
	valuesMulti := []interface{}{
		bson.D{
			{Key: "key_1", Value: "value_1"},
			{Key: "key_2", Value: "value_2"},
		},
		bson.D{
			{"tags", bson.A{"tag_1", "tag_2", "tag_3"}},
		},
	}

	data2, _ := mongoEpisodes.InsertMany(timeOut, valuesMulti)

	fmt.Println(data2.InsertedIDs...)
}
