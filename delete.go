package main

import (
    "context"
    "fmt"
	"log"
    "go.mongodb.org/mongo-driver/bson"
)

func delete() {
	client := connect()    
		
	collection := client.Database("DbTest").Collection("CollectionTest")

	filter := bson.D {{"sobrenome", "Tres"}}

	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	disconnect(client)
}