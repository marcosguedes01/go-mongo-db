package main

import (
    "context"
    "fmt"
    "log"
)

func write() {
	client := connect()    
		
	collection := client.Database("DbTest").Collection("CollectionTest")

	item := Person { "Nome Teste 2", "Sobrenome Teste 2" }

	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

	disconnect(client)
}