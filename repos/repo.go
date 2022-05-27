package repo

import (
	"context"
	"time"

	database "github.com/Vicjocaso/Go-mongodb/Database"
	"github.com/Vicjocaso/Go-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Collection = database.GetCollection("user")
var ctx = context.Background()

func Create(user models.User) error {

	var err error

	_, err = Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (models.Users, error) {

	var users models.Users

	filter := bson.D{}

	cur, err := Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user models.User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func Update(user models.User, userId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}

	Update := bson.M{
		"$set": bson.M{

			"name":    user.Name,
			"email":   user.Email,
			"updated": time.Now(),
		},
	}

	_, err = Collection.UpdateOne(ctx, filter, Update)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {

	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
