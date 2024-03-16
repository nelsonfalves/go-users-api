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

func Get() (canonical.User, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:password@localhost:27017"))

	if err != nil {
		return canonical.User{}, err
	}

	collection := client.Database("user-database").Collection("users")

	var user canonical.User
	res, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return canonical.User{}, err
	}

	for res.Next(context.Background()) {
		err := res.Decode(&user)
		if err != nil {
			return canonical.User{}, err
		}
	}

	fmt.Println(user)

	return user, nil
}
