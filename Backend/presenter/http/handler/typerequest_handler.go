package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type TyperequestHandler interface {
	GetTyperequests(c echo.Context) error
	CreateTyperequest(c echo.Context) error
	UpdateTyperequest(c echo.Context) error
	DeleteTyperequest(c echo.Context) error
}

type typerequestHandler struct {
	TyperequestUseCase usecase.TyperequestUseCase
}

func NewTyperequestHandler(u usecase.TyperequestUseCase) TyperequestHandler {
	return &typerequestHandler{u}
}

func (h *typerequestHandler) GetTyperequests(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.TyperequestUseCase.GetTyperequests(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Typerequests does not exist.")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *typerequestHandler) CreateTyperequest(c echo.Context) error {
	object := &model.Typerequest{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Typerequestname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.TyperequestUseCase.CreateTyperequest(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *typerequestHandler) UpdateTyperequest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}
	object := &model.Typerequest{
		Typerequestid: id,
	}
	if err := c.Bind(object); err != nil {
		return err
	}
	if object.Typerequestid == 0 || object.Typerequestname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.TyperequestUseCase.UpdateTyperequest(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *typerequestHandler) DeleteTyperequest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.TyperequestUseCase.DeleteTyperequest(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
