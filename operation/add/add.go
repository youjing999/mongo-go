package main

import (
	"context"
	"fmt"
	"mongo-go/global"
	"mongo-go/initMongo"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	initMongo.InitMongoDB()
	insertMany()
}

func insertOne(s Student) {
	collection := global.Client.Database("go_db").Collection("student")
	insertRes, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		panic(err)
	}

	fmt.Println("insert a record", insertRes.InsertedID)
}

func insertMany() {
	var students []interface{}
	students = append(students, Student{Name: "yj", Age: 19})
	students = append(students, Student{Name: "calvin", Age: 20})
	collection := global.Client.Database("go_db").Collection("student")
	manyRes, err := collection.InsertMany(context.Background(), students)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted multiple documents: ", manyRes.InsertedIDs)
}
