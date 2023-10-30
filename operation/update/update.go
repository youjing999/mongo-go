package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"mongo-go/global"
	"mongo-go/initMongo"
	"time"
)

type Student struct {
	Name string
	Age  int
}

func update() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := global.Client.Database("go_db").Collection("student")

	update := bson.D{{"$set", bson.D{{"Name", "xzq"}, {"Age", 22}}}}

	updateRes, err := collection.UpdateMany(ctx, bson.D{{"name", "yj"}}, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateRes.ModifiedCount)
}

func main() {
	initMongo.InitMongoDB()
	update()
}
