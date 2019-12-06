package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
)

func disconnect(client *mongo.Client) {
    //client := connect()    

    err := client.Disconnect(context.TODO())

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Connection to MongoDB closed.")
    }
}