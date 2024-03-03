package repositories

import (
	"context"
	"fmt"
	"go-quickstart/internal/canonical"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Create() error {
	variavel := "teste"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://nelsonalves117:083254sp@cluster0.xn2hv6l.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		return err
	}

	collection := client.Database("user-database").Collection("users")

	res, err := collection.InsertOne(context.Background(), canonical.User{
		Name: variavel,
		Id:   "1",
	})

	if err != nil {
		return err
	}

	id := res.InsertedID

	fmt.Println(id)

	return nil
}

func Get(id string) (canonical.User, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://nelsonalves117:083254sp@cluster0.xn2hv6l.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		return canonical.User{}, err
	}

	collection := client.Database("user-database").Collection("users")

	objID, _ := primitive.ObjectIDFromHex(id)
	result := collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}})
	var user canonical.User
	err = result.Decode(&user)
	if err != nil {
		return canonical.User{}, err
	}

	return user, nil
}
