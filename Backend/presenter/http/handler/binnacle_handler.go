package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type BinnacleHandler interface {
	CreateBinnacleByID(c echo.Context) error
	FinishBinnacleByID(c echo.Context) error
	GetBinnacleById(c echo.Context) error
	GetBinnacleByIdMovil(c echo.Context) error
	CheckIn(c echo.Context) error
	CheckOut(c echo.Context) error
	MealDay(c echo.Context) error
	UpdateDay(c echo.Context) error
}

type binnacleHandler struct {
	usecase.BinnacleUseCase
}

func NewBinnacleHandler(u usecase.BinnacleUseCase) BinnacleHandler {
	return &binnacleHandler{u}
}

func (h *binnacleHandler) CreateBinnacleByID(c echo.Context) error {
	object := &model.CreateBinnacle{}
	if err := c.Bind(object); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if object.Userid == 0 || object.Month == 0 || object.Year == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	err := h.BinnacleUseCase.CreateBinnacleByID(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Binnacle can not Create.")
	}

	return c.JSON(http.StatusOK, true)
}

func (h *binnacleHandler) FinishBinnacleByID(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

func (h *binnacleHandler) GetBinnacleById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Client ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	result, err := h.BinnacleUseCase.GetBinnacleById(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *binnacleHandler) GetBinnacleByIdMovil(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Client ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	result, err := h.BinnacleUseCase.GetBinnacleByIdMovil(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *binnacleHandler) CheckIn(c echo.Context) error {
	object := &model.Checktime{}
	if err := c.Bind(object); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	fmt.Println(object)
	if object.Userid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	err := h.BinnacleUseCase.CheckIn(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, true)
}

func (h *binnacleHandler) CheckOut(c echo.Context) error {
	object := &model.Departuretime{}
	if err := c.Bind(object); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if object.Userid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	err := h.BinnacleUseCase.CheckOut(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, true)
}

func (h *binnacleHandler) MealDay(c echo.Context) error {
	object := &model.Mealtime{}
	if err := c.Bind(object); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if object.Userid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	err := h.BinnacleUseCase.MealDay(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, true)
}

func (h *binnacleHandler) UpdateDay(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}
