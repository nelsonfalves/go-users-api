package service

import (
	"fmt"
	"go-quickstart/internal/canonical"
	"go-quickstart/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Create(user canonical.User) error
	Get() ([]canonical.User, error)
	GetById(id string) (canonical.User, error)
	Update(id string, user canonical.User) error
	Delete(id string) error
}

type service struct {
	repo repositories.Repository
}

func New() Service {
	return &service{
		repo: repositories.New(),
	}
}

func (service *service) Create(user canonical.User) error {
	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()

	err := service.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *service) Get() ([]canonical.User, error) {
	user, err := service.repo.Get()
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (service *service) GetById(id string) (canonical.User, error) {
	user, err := service.repo.GetById(id)
	if err != nil {
		return canonical.User{}, err
	}

	return user, nil

}

func (service *service) Update(id string, user canonical.User) error {
	err := service.repo.Update(id, user)
	if err != nil {
		return err
	}

	return nil

}

func (service *service) Delete(id string) error {
	user, err := service.repo.GetById(id)
	if err != nil {
		return err
	}

	if user.Id == "" {
		return fmt.Errorf("user not found on db")
	}

	err = service.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
