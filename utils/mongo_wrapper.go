package utils

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectToMongoDB() {
    // MongoDB connection URI
    uri := "mongodb://localhost:27017" // replace with your MongoDB URI

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    // Test the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
}

func DisconnectMongoDB() {
    if err := client.Disconnect(context.TODO()); err != nil {
        log.Fatal(err)
    }
    log.Println("Disconnected from MongoDB.")
}