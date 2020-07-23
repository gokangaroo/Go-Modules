package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func main() {
	var (
		ctx        = context.Background()
		test       = &Test{Name: "test"}
		res        *mongo.InsertOneResult
		collection *mongo.Collection
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://huijia:123456@localhost:27017/?replicaSet=replicaset&authSource=demo"))
	if err != nil {
		goto ERR
	}
	err = client.Ping(ctx, readpref.Primary()) //read prefer
	if err != nil {
		goto ERR
	}
	// todo is this reasonable
	test.ID = primitive.NewObjectID()
	collection = client.Database("demo").Collection("test")
	res, err = collection.InsertOne(ctx, test)
	if err != nil {
		goto ERR
	}
	log.Println(res)
	return
ERR:
	log.Fatal(err)
}

type Test struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
