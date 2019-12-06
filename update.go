package main

import (
    "context"
    "fmt"
	"log"
    "go.mongodb.org/mongo-driver/bson"
)

func update() {
	client := connect()    
		
	collection := client.Database("DbTest").Collection("CollectionTest")

	filter := bson.D {{"sobrenome", "Um"}}
    update := bson.D{
        {"$set", bson.D {
			{"nome", "Teste 1"},
			{"sobrenome", "Teste Um"},
        }},
    }

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	disconnect(client)
}