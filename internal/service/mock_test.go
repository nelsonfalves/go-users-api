package service

import (
	"github.com/nelsonalves117/go-users-api/internal/canonical"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllUsers() ([]canonical.User, error) {
	args := m.Called()
	return args.Get(0).([]canonical.User), args.Error(1)
}

func (m *MockRepository) GetUserById(id string) (canonical.User, error) {
	args := m.Called(id)
	return args.Get(0).(canonical.User), args.Error(1)
}

func (m *MockRepository) CreateUser(user canonical.User) (canonical.User, error) {
	args := m.Called(user)
	return args.Get(0).(canonical.User), args.Error(1)
}

func (m *MockRepository) UpdateUser(id string, user canonical.User) (canonical.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(canonical.User), args.Error(1)
}

func (m *MockRepository) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
