package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"

	"github.com/grupogit/RMNGR002001/Backend/usecase"

	"github.com/labstack/echo/v4"
)

type SoftskillHandler interface {
	GetSoftskills(c echo.Context) error
	CreateSoftskill(c echo.Context) error
	UpdateSoftskill(c echo.Context) error
	DeleteSoftskill(c echo.Context) error
}

type softskillHandler struct {
	SoftskillUseCase usecase.SoftskillUseCase
}

func NewSoftskillHandler(u usecase.SoftskillUseCase) SoftskillHandler {
	return &softskillHandler{u}
}

func (h *softskillHandler) GetSoftskills(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.SoftskillUseCase.GetSoftskills(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Softskills does not exist.")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *softskillHandler) CreateSoftskill(c echo.Context) error {
	skill := &model.Softskill{}
	if err := c.Bind(skill); err != nil {
		return err
	}

	if skill.Softskillname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.SoftskillUseCase.CreateSoftskill(ctx, skill)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *softskillHandler) UpdateSoftskill(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	skill := &model.Softskill{
		Softskillid: id,
	}
	if err := c.Bind(&skill); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if skill.Softskillid == 0 || skill.Softskillname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	result, err := h.SoftskillUseCase.UpdateSoftskill(ctx, skill)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *softskillHandler) DeleteSoftskill(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.SoftskillUseCase.DeleteSoftskill(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
