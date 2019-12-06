package main

import (
    "context"
    "fmt"
	"log"
    "go.mongodb.org/mongo-driver/bson"
)

func find() {
	client := connect()    
		
	collection := client.Database("DbTest").Collection("CollectionTest")

	var result Person
	filter := bson.D {{"sobrenome", "Dois"}}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	disconnect(client)
}