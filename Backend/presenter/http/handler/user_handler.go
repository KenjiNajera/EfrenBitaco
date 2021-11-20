package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/helper"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	ActivateUser(c echo.Context) error
	DesactivateUser(c echo.Context) error
	GetUsers(c echo.Context) error
	// GetMyData(c echo.Context) error
	GetMyDataMovil(c echo.Context) error
	// UpdateMyData(c echo.Context) error
}

type userHandler struct {
	UserUscase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	image, _ := c.FormFile("file")
	data := c.FormValue("data")

	object := &model.User{}
	err := json.Unmarshal([]byte(data), object)
	if err != nil {
		fmt.Println(err)
	}

	var str string
	if image != nil {
		str, err = helper.Upload(image)
		object.Resourcedata.Photouser = str
		if err != nil {
			return err
		}
	}

	if object.Email == "" ||
		object.Roleid == 0 ||
		object.Resourcedata.Firstname == "" ||
		object.Resourcedata.Lastname == "" ||
		object.Resourcedata.Personalemail == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.UserUscase.CreateUser(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	image, _ := c.FormFile("file")
	data := c.FormValue("data")

	object := &model.User{}
	err := json.Unmarshal([]byte(data), object)
	if err != nil {
		return err
	}

	if image != nil {
		object.Resourcedata.Photouser, err = helper.Upload(image)
		fmt.Println("entre a la modificacion de la imagen")
		if err != nil {
			return err
		}
	}

	fmt.Println(object)

	if object.Email == "" ||
		object.Roleid == 0 ||
		object.Resourcedata.Firstname == "" ||
		object.Resourcedata.Lastname == "" ||
		object.Resourcedata.Personalemail == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.UserUscase.UpdateUser(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(result)
	return c.JSON(http.StatusCreated, result)
}

func (h *userHandler) ActivateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.UserUscase.ActivateUser(ctx, id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *userHandler) DesactivateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.UserUscase.DesactivateUser(ctx, id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *userHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.UserUscase.GetUsers(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *userHandler) GetMyDataMovil(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.UserUscase.GetMyDataMovil(ctx, id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
