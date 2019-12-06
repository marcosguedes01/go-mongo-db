package main

import (
    "context"
    "fmt"
    "log"
)

func writeMany() {
	client := connect()    
		
	collection := client.Database("DbTest").Collection("CollectionTest")

	teste1 := Person{"Teste", "Um"}
    teste2 := Person{"Teste", "Dois"}
	teste3 := Person{"Teste", "Tres"}
	
	testes := []interface{}{teste1, teste2, teste3}

	insertManyResult, err := collection.InsertMany(context.TODO(), testes)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	disconnect(client)
}