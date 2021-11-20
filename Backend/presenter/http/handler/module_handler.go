package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type ModuleHandler interface {
	GetModules(c echo.Context) error
	CreateModule(c echo.Context) error
	UpdateModule(c echo.Context) error
	DeleteModule(c echo.Context) error
}

type moduleHandler struct {
	ModuleUseCase usecase.ModuleUseCase
}

func NewModuleHandler(u usecase.ModuleUseCase) ModuleHandler {
	return &moduleHandler{u}
}

func (h *moduleHandler) GetModules(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ModuleUseCase.GetModules(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Modules does not exist.")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *moduleHandler) CreateModule(c echo.Context) error {
	object := &model.Module{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Modulename == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ModuleUseCase.CreateModule(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *moduleHandler) UpdateModule(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}
	object := &model.Module{
		Moduleid: id,
	}
	if err := c.Bind(object); err != nil {
		return err
	}
	if object.Moduleid == 0 || object.Modulename == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ModuleUseCase.UpdateModule(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *moduleHandler) DeleteModule(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.ModuleUseCase.DeleteModule(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
