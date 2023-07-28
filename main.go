package main

import (
	"context"
	"fmt"
	"go_tutorial/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)  
type Stock struct {
	StockID int64  `json:"stockid" validate:"required"`
	Name    *string `json:"name" validate:"required,min=3,max=100"`
	Price   int64  `json:"price" validate:"required,min=1"`
	Company string `json:"company" validate:"required"`
}
func main() {
	stockName := "e"
	stock := Stock{
		StockID: 12345,             // StockID is a valid non-zero value
		Name:    &stockName,          // Name is valid and has a length of 6 characters
     // Company is a valid non-empty string
	}

	stockMap := bson.M{
		"stockid": stock.StockID,
		"name":    stock.Name,
	
	
	}
	data := bson.D{}
	for key, value := range stockMap {
		data = append(data, bson.E{Key: key, Value: value})
	}
	var client = database.DBinstance()
	var collection = database.OpenCollection(client, "test")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resultInsertionNumber, _ := collection.InsertOne(ctx, data)
	fmt.Println("Inserted document ID:", resultInsertionNumber.InsertedID)
	fmt.Println(" stock", stock)
}