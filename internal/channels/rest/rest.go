package rest

import (
	"errors"
	"go-quickstart/internal/config"
	"go-quickstart/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
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

	router.GET("/", rest.GetAllUsers)
	router.GET("/user/:id", rest.GetUserById)
	router.POST("/create", rest.CreateUser)
	router.PUT("/update/:id", rest.UpdateUser)
	router.DELETE("/delete/:id", rest.DeleteUser)

	return router.Start(":" + config.Get().Port)
}

func (rest *rest) CreateUser(c echo.Context) error {
	var user user

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	err = rest.service.CreateUser(toCanonical(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusNoContent, nil)
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

func (rest *rest) UpdateUser(c echo.Context) error {
	var user user

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	id := c.Param("id")
	err = rest.service.UpdateUser(id, toCanonical(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, user)
}

func (rest *rest) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := rest.service.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, nil)
}
