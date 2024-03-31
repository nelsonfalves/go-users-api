package repositories

import (
	"context"
	"fmt"
	"go-quickstart/internal/canonical"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
}

type repository struct {
}

func New() Repository {
	return &repository{}
}

func Create(user canonical.User) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		return err
	}

	collection := client.Database("user-database").Collection("users")

	res, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		return err
	}

	id := res.InsertedID

	fmt.Println(id)

	return nil
}

func Get() ([]canonical.User, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		return nil, err
	}

	collection := client.Database("user-database").Collection("users")

	var users []canonical.User
	res, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for res.Next(context.Background()) {
		var user canonical.User
		err := res.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetById(id string) (canonical.User, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:password@localhost:27017"))

	if err != nil {
		return canonical.User{}, err
	}

	collection := client.Database("user-database").Collection("users")

	result := collection.FindOne(context.Background(), bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	})

	var user canonical.User

	result.Decode(&user)

	err = result.Decode(&user)
	if err != nil {
		return canonical.User{}, err
	}

	return user, nil
}

func Update(id string, user canonical.User) error {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:password@localhost:27017"))

	if err != nil {
		return err
	}

	collection := client.Database("user-database").Collection("users")

	filter := bson.D{{Key: "_id", Value: id}}
	fields := bson.M{"$set": user}

	_, err = collection.UpdateOne(context.Background(), filter, fields)

	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		return err
	}

	collection := client.Database("user-database").Collection("users")

	_, err = collection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	return nil

}
