package initMongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongo-go/global"
)

func InitMongoDB() {
	//设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//连接MongoDB
	var err error
	global.Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = global.Client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
}
