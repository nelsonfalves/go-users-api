package rest

import (
	"errors"
	"net/http"

	"github.com/nelsonalves117/go-users-api/internal/config"
	"github.com/nelsonalves117/go-users-api/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Rest interface {
	Start() error
}

type rest struct {
	service service.Service
}

func New() Rest {
	return &rest{
		service: service.New(),
	}
}

func (rest *rest) Start() error {
	router := echo.New()

	router.Use(middleware.Logger())

	router.GET("/users", rest.GetAllUsers)
	router.GET("/users/:id", rest.GetUserById)
	router.POST("/users/create", rest.CreateUser)
	router.PUT("/users/update/:id", rest.UpdateUser)
	router.DELETE("/users/delete/:id", rest.DeleteUser)

	return router.Start(":" + config.Get().Port)
}

func (rest *rest) GetAllUsers(c echo.Context) error {
	userSlice, err := rest.service.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, userSlice)
}

func (rest *rest) GetUserById(c echo.Context) error {
	id := c.Param("id")

	user, err := rest.service.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, user)
}

func (rest *rest) CreateUser(c echo.Context) error {
	var user userRequest

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	createdUser, err := rest.service.CreateUser(toCanonical(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusCreated, toResponse(createdUser))
}

func (rest *rest) UpdateUser(c echo.Context) error {
	var user userRequest

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	id := c.Param("id")
	updatedUser, err := rest.service.UpdateUser(id, toCanonical(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, toResponse(updatedUser))
}

func (rest *rest) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := rest.service.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, nil)
}
