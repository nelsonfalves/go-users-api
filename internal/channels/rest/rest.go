package rest

import (
	"errors"
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

	router.GET("/", rest.Get)
	router.GET("/user/:id", rest.GetById)
	router.POST("/create", rest.Create)
	router.PUT("/update/:id", rest.Update)
	router.DELETE("/delete/:id", rest.Delete)

	return router.Start(":8080")
}

func (rest *rest) Create(c echo.Context) error {
	var user user

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	err = rest.service.Create(toCanonical(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (rest *rest) Get(c echo.Context) error {
	users, err := rest.service.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, users)
}

func (rest *rest) GetById(c echo.Context) error {
	id := c.Param("id")
	user, err := rest.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, user)
}

func (rest *rest) Update(c echo.Context) error {
	var user user

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	id := c.Param("id")
	err = rest.service.Update(id, toCanonical(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, user)
}

func (rest *rest) Delete(c echo.Context) error {
	id := c.Param("id")
	err := rest.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, nil)
}
