package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = "Cluster0"

func GetCollection(conllection string) *mongo.Collection {

	uri := "Uri MongoConnection string "

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection((conllection))

}
