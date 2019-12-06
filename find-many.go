package main

import (
    "context"
    "fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findMany() {
	client := connect()    
		
	collection := client.Database("DbTest").Collection("CollectionTest")

	findOptions := options.Find()
	//findOptions.SetLimit(2)

	var results []*Person
	filter := bson.D {}
	
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Person
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Item: ", elem.Nome + " " + elem.Sobrenome + "\n")

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	
	disconnect(client)
}