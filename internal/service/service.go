package service

import (
	"fmt"
	"time"

	"github.com/nelsonalves117/go-users-api/internal/canonical"
	"github.com/nelsonalves117/go-users-api/internal/repositories"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Service interface {
	GetAllUsers() ([]canonical.User, error)
	GetUserById(id string) (canonical.User, error)
	CreateUser(user canonical.User) (canonical.User, error)
	UpdateUser(id string, user canonical.User) (canonical.User, error)
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

func (service *service) GetAllUsers() ([]canonical.User, error) {
	user, err := service.repo.GetAllUsers()
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to get all users")
		return nil, err
	}

	return user, nil
}

func (service *service) GetUserById(id string) (canonical.User, error) {
	user, err := service.repo.GetUserById(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to get a user")
		return canonical.User{}, err
	}

	return user, nil
}

func (service *service) CreateUser(user canonical.User) (canonical.User, error) {
	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()

	user, err := service.repo.CreateUser(user)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to create a user")
		return canonical.User{}, err
	}

	return user, nil
}

func (service *service) UpdateUser(id string, user canonical.User) (canonical.User, error) {
	user, err := service.repo.UpdateUser(id, user)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to update a user")
		return canonical.User{}, err
	}

	return user, nil
}

func (service *service) DeleteUser(id string) error {
	user, err := service.repo.GetUserById(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to get a user")
		return err
	}

	if user.Id == "" {
		return fmt.Errorf("user not found on db")
	}

	err = service.repo.DeleteUser(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to delete a user")
		return err
	}

	return nil
}
