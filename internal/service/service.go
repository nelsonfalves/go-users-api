package service

import (
	"fmt"
	"go-quickstart/internal/canonical"
	"go-quickstart/internal/repositories"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateUser(user canonical.User) (canonical.User, error)
	GetAllUsers() ([]canonical.User, error)
	GetUserById(id string) (canonical.User, error)
	UpdateUser(id string, user canonical.User) error
	DeleteUser(id string) error
}

type service struct {
	repo repositories.Repository
}

func New() Service {
	return &service{
		repo: repositories.New(),
	}
}

func (service *service) CreateUser(user canonical.User) (canonical.User, error) {
	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()

	err := service.repo.CreateUser(user)
	if err != nil {
		logrus.WithError(err).Error("Error occurred when trying to create an user.")
		return canonical.User{}, err
	}

	return user, nil
}

func (service *service) GetAllUsers() ([]canonical.User, error) {
	user, err := service.repo.GetAllUsers()
	if err != nil {
		logrus.WithError(err).Error("Error occurred when trying to get all users.")
		return nil, err
	}

	return user, nil

}

func (service *service) GetUserById(id string) (canonical.User, error) {
	user, err := service.repo.GetUserById(id)
	if err != nil {
		logrus.WithError(err).Error("Error occurred when trying to get an user.")
		return canonical.User{}, err
	}

	return user, nil

}

func (service *service) UpdateUser(id string, user canonical.User) error {
	err := service.repo.UpdateUser(id, user)
	if err != nil {
		logrus.WithError(err).Error("Error occurred when trying to update an user.")
		return err
	}

	return nil

}

func (service *service) DeleteUser(id string) error {
	user, err := service.repo.GetUserById(id)
	if err != nil {
		logrus.WithError(err).Error("Error occurred when trying to get an user.")
		return err
	}

	if user.Id == "" {
		return fmt.Errorf("user not found on db")
	}

	err = service.repo.DeleteUser(id)
	if err != nil {
		logrus.WithError(err).Error("Error occurred when trying to delete an user.")
		return err
	}

	return nil
}
