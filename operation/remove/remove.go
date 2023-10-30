package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"mongo-go/global"
	"mongo-go/initMongo"
)

func remove() {
	c := global.Client.Database("go_db").Collection("Student")

	dr, err := c.DeleteMany(context.Background(), bson.D{{"Name", "calvin"}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", dr.DeletedCount)
}
func main() {
	initMongo.InitMongoDB()
	remove()
}
