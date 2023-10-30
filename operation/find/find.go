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

func find() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := global.Client.Database("go_db").Collection("student")

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	var s Student
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)

		////Deprecated
		//fmt.Printf("result.Map(): %v\n", result.Map()["name"])
		marshalByte, err := bson.Marshal(result)
		err = bson.Unmarshal(marshalByte, &s)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func findWithAnd() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := global.Client.Database("go_db").Collection("student")

	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"name", "tim"}},
			bson.D{{"age", 20}},
		}},
	}

	var students []Student
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var student Student
		err := cur.Decode(&student)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, student)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("results: %v\n", students)
}

func findWithOr() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := global.Client.Database("go_db").Collection("student")

	filter := bson.M{
		"$or": []bson.M{
			bson.M{"name": "tim"},
			bson.M{"age": 20},
		},
	}

	var students []Student
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var student Student
		err := cur.Decode(&student)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, student)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("results: %v\n", students)
}

func main() {
	initMongo.InitMongoDB()
	find()
}
