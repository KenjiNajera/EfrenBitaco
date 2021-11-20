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

type ClientHandler interface {
	GetClients(c echo.Context) error
	GetClientByID(c echo.Context) error
	CreateClient(c echo.Context) error
	UpdateClient(c echo.Context) error
	DeleteClient(c echo.Context) error
	Reactivate(c echo.Context) error
}

type clientHandler struct {
	ClientUseCase usecase.ClientUseCase
}

func NewClientHandler(u usecase.ClientUseCase) ClientHandler {
	return &clientHandler{u}
}

func (h *clientHandler) GetClients(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ClientUseCase.GetClients(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *clientHandler) GetClientByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Client ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ClientUseCase.GetClientByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *clientHandler) CreateClient(c echo.Context) error {
	object := &model.Client{}
	image, _ := c.FormFile("file")
	data := c.FormValue("data")
	err := json.Unmarshal([]byte(data), object)
	if err != nil {
		fmt.Println(err)
	}

	if image != nil {
		object.Photoclient, err = helper.Upload(image)
		if err != nil {
			return err
		}
	}

	if object.Clientname == "" || object.Socialreason == "" || object.Rfc == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ClientUseCase.CreateClient(ctx, object)
	client := &model.Clients{
		Clientid:    result,
		Clientname:  object.Clientname,
		Photoclient: object.Photoclient,
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, client)
	}

	return c.JSON(http.StatusOK, client)
}

func (h *clientHandler) UpdateClient(c echo.Context) error {
	object := &model.Client{}
	image, _ := c.FormFile("file")
	data := c.FormValue("data")
	err := json.Unmarshal([]byte(data), object)
	if err != nil {
		fmt.Println(err)
	}

	var str string
	if image != nil {
		str, err = helper.Upload(image)
		object.Photoclient = str
		if err != nil {
			return err
		}
	}

	if object.Clientname == "" || object.Socialreason == "" || object.Rfc == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ClientUseCase.UpdateClient(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Client can not Create.")
	}
	return c.JSON(http.StatusOK, result)
}

func (h *clientHandler) DeleteClient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Client ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.ClientUseCase.DeleteClient(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}

func (h *clientHandler) Reactivate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Client ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ClientUseCase.Reactivate(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}
