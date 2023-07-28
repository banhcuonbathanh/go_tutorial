package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Client *mongo.Client
func DBinstance() *mongo.Client {
	fmt.Println(" this is DBinstance function ")
	clientOptions := options.Client().ApplyURI("mongodb+srv://vu:ez7zOr4f6jjVA6Zg@cluster0.exj1kwg.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("testdatabase").Collection(collectionName)

	return collection

}