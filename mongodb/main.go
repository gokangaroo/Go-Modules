package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	// 0. replica
	// MONGODB_URI="mongodb://localhost:27017,localhost:27018,localhost:27018/?replicaSet=rs1"
	// 1.create a client
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	//printErr(err)
	// 2. connect to server
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//err = client.Connect(ctx)
	//printErr(err)
	// 3.one step do 1-2
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	printErr(err)
	err = client.Ping(ctx, readpref.Primary()) //read prefer
	// 4.get a collection & find
	var (
		r       bson.M
		res     *mongo.Cursor
		marshal []byte
	)
	collection := client.Database("demo").Collection("user")
	// 参数使用bson.M最清晰, 直接是个map, bson.D的话有点麻烦了就, 结构体.
	//filter := bson.M{
	//	"userId": 470580,
	//}
	//res, err = collection.Find(ctx, filter)
	//printErr(err)
	//decodeOne(ctx, res, &r)
	//marshal, _ = json.Marshal(r)
	//fmt.Println("result1: ", string(marshal))
	// 5. pipeline
	groupStage := []bson.M{
		{
			"$match": bson.M{
				"userId": 470580,
			}}, {
			"$lookup": bson.M{
				"from":         "contract",
				"localField":   "contractId",
				"foreignField": "contractId",
				"as":           "contract",
			}}, {
			"$unwind": bson.M{
				"path":                       "$contract",
				"preserveNullAndEmptyArrays": true,
			}}, {
			"$unwind": bson.M{
				"path":                       "$components",
				"preserveNullAndEmptyArrays": true,
			}}, {
			"$match": bson.M{
				"components.name": "summit",
			}}, {
			"$unwind": bson.M{
				"path":                       "$contract.components",
				"preserveNullAndEmptyArrays": true,
			}}, {
			"$match": bson.M{
				"contract.components.name": "summit",
			}}, {
			// project没有生效, 因为是0的缘故?
			// 原来是,上面单独查的时候就已经有了, 所以虽然这里pipeline没有, 但是实际已经存在了.
			// TODO 业务可以先find contract, 再find user来直接覆盖.
			"$project": bson.M{
				"_id":        0,
				"contractId": 0,
			}},
	}
	opts := options.Aggregate().SetMaxTime(2 * time.Second)
	res, err = collection.Aggregate(ctx, groupStage, opts)
	printErr(err)
	decodeOne(ctx, res, &r)
	marshal, _ = json.Marshal(r)
	fmt.Println("result2: ", string(marshal))
}

func printErr(err error) {
	if err != nil {
		panic(err)
	}
}

// find是cur, findOne可以直接decode.
func decodeOne(ctx context.Context, cur *mongo.Cursor, r *bson.M) {
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		err := cur.Decode(r)
		printErr(err)
		return
	}
}
