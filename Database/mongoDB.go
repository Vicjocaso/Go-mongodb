package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = "Cluster0"

func GetCollection(conllection string) *mongo.Collection {

	uri := "mongodb+srv://Vicjocaso:Admin123@cluster0.ghnre.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection((conllection))

}
