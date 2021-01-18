// 测试 mongodb ACID 事务功能
// 参考网址: https://developer.mongodb.com/quickstart/golang-multi-document-acid-transactions
// @author moqi
// On 2021/1/18 22:14:31
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// Episode represents the schema for the "Episodes" collection
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27018"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	database := client.Database("quickstart")
	episodesCollection := database.Collection("episodes")

	database.RunCommand(context.TODO(), bson.D{{"create", "episodes"}})

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		result, err := episodesCollection.InsertOne(
			sessionContext,
			Episode{
				Title:    "A Transaction Episode for the Ages",
				Duration: 15,
			},
		)

		// 模拟第一次执行后失败
		// return nil, errors.New("insert first record not success")

		if err != nil {
			return nil, err
		}

		fmt.Println("the first record id: ", result.InsertedID)

		result, err = episodesCollection.InsertOne(
			sessionContext,
			Episode{
				Title:    "Transactions for All",
				Duration: 2,
			},
		)

		if err != nil {
			return nil, err
		}

		fmt.Println("the second record id: ", result.InsertedID)
		return result, err
	}
	_, err = session.WithTransaction(context.Background(), callback, txnOpts)
	if err != nil {
		panic(err)
	}
}
