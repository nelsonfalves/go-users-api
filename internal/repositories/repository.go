package repositories

import (
	"context"
	"go-quickstart/internal/canonical"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(user canonical.User) error
	Get() ([]canonical.User, error)
	GetById(id string) (canonical.User, error)
	Update(id string, user canonical.User) error
	Delete(id string) error
}

type repository struct {
	collection *mongo.Collection
}

func New() Repository {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		panic(err)
	}

	return &repository{
		collection: client.Database("user-database").Collection("users"),
	}
}

func (repo *repository) Create(user canonical.User) error {
	_, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Get() ([]canonical.User, error) {
	var users []canonical.User

	res, err := repo.collection.Find(context.Background(), bson.D{})
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

func (repo *repository) GetById(id string) (canonical.User, error) {
	var user canonical.User

	err := repo.collection.FindOne(context.Background(), bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}).Decode(&user)

	if err != nil {
		return canonical.User{}, err
	}

	return user, nil
}

func (repo *repository) Update(id string, user canonical.User) error {
	filter := bson.D{{Key: "_id", Value: id}}
	fields := bson.M{"$set": user}

	_, err := repo.collection.UpdateOne(context.Background(), filter, fields)

	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Delete(id string) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}

	return nil
}
