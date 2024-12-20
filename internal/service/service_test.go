package service

import (
	"errors"
	"testing"
	"time"

	"github.com/nelsonalves117/go-users-api/internal/canonical"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsers_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	usersTest := []canonical.User{
		{
			Id:        "xpto",
			Name:      "test",
			Email:     "test@email.com",
			Password:  "xpto",
			CreatedAt: time.Now(),
		},
	}

	mockRepo.On("GetAllUsers").Return(usersTest, nil)

	service := &service{
		repo: mockRepo,
	}

	users, err := service.GetAllUsers()

	assert.Nil(t, err)
	assert.Equal(t, "xpto", users[0].Id)
	assert.Equal(t, "test", users[0].Name)
	assert.Equal(t, "test@email.com", users[0].Email)
	assert.Equal(t, "xpto", users[0].Password)
	assert.True(t, users[0].CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetAllUsers").Return([]canonical.User{}, errors.New("error occurred while trying to get all users"))

	service := &service{
		repo: mockRepo,
	}

	users, err := service.GetAllUsers()

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Equal(t, "error occurred while trying to get all users", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetUserById_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	userTest := canonical.User{
		Id:        "xpto",
		Name:      "test",
		Email:     "test@email.com",
		Password:  "xpto",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetUserById", "xpto").Return(userTest, nil)

	service := &service{
		repo: mockRepo,
	}

	user, err := service.GetUserById("xpto")

	assert.Nil(t, err)
	assert.Equal(t, "xpto", user.Id)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@email.com", user.Email)
	assert.Equal(t, "xpto", user.Password)
	assert.True(t, user.CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}

func TestGetUserById_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetUserById", "xpto").Return(canonical.User{}, errors.New("error occurred while trying to get a user"))

	service := &service{
		repo: mockRepo,
	}

	user, err := service.GetUserById("xpto")

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Equal(t, "error occurred while trying to get a user", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	userTest := canonical.User{
		Name:     "test",
		Email:    "test@email.com",
		Password: "xpto",
	}

	mockRepo.On("CreateUser", mock.MatchedBy(func(user canonical.User) bool {
		return user.Name == "test" && user.Email == "test@email.com" && user.Password == "xpto"
	})).Return(userTest, nil)

	service := &service{
		repo: mockRepo,
	}

	user, err := service.CreateUser(userTest)

	assert.Nil(t, err)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@email.com", user.Email)
	assert.Equal(t, "xpto", user.Password)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	userTest := canonical.User{
		Name:     "test",
		Email:    "test@email.com",
		Password: "xpto",
	}

	mockRepo.On("CreateUser", mock.MatchedBy(func(user canonical.User) bool {
		return user.Name == "test" && user.Email == "test@email.com" && user.Password == "xpto"
	})).Return(canonical.User{}, errors.New("error occurred while trying to create a user"))

	service := &service{
		repo: mockRepo,
	}

	user, err := service.CreateUser(userTest)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Equal(t, "error occurred while trying to create a user", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	userTest := canonical.User{
		Name:     "test",
		Email:    "test@email.com",
		Password: "xpto",
	}

	updatedUser := canonical.User{
		Id:        "xpto",
		Name:      "test",
		Email:     "test@email.com",
		Password:  "xpto",
		CreatedAt: time.Now(),
	}

	mockRepo.On("UpdateUser", "xpto", userTest).Return(updatedUser, nil)

	service := &service{
		repo: mockRepo,
	}

	user, err := service.UpdateUser("xpto", userTest)

	assert.Nil(t, err)
	assert.Equal(t, "xpto", user.Id)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@email.com", user.Email)
	assert.Equal(t, "xpto", user.Password)
	assert.True(t, user.CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	userTest := canonical.User{
		Name:     "test",
		Email:    "test@email.com",
		Password: "xpto",
	}

	mockRepo.On("UpdateUser", "xpto", userTest).Return(canonical.User{}, errors.New("error occurred while trying to update a user"))

	service := &service{
		repo: mockRepo,
	}

	user, err := service.UpdateUser("xpto", userTest)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Equal(t, "error occurred while trying to update a user", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	userTest := canonical.User{
		Id:        "xpto",
		Name:      "test",
		Email:     "test@email.com",
		Password:  "xpto",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetUserById", "xpto").Return(userTest, nil)

	mockRepo.On("DeleteUser", "xpto").Return(nil)

	service := &service{
		repo: mockRepo,
	}

	err := service.DeleteUser("xpto")

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	userTest := canonical.User{
		Id:        "xpto",
		Name:      "test",
		Email:     "test@email.com",
		Password:  "xpto",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetUserById", "xpto").Return(userTest, nil)

	mockRepo.On("DeleteUser", "xpto").Return(errors.New("error occurred while trying to delete a user"))

	service := &service{
		repo: mockRepo,
	}

	err := service.DeleteUser("xpto")

	assert.NotNil(t, err)
	assert.Equal(t, "error occurred while trying to delete a user", err.Error())

	mockRepo.AssertExpectations(t)
}
