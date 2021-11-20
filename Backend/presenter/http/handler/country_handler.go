package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type CountryHandler interface {
	GetCountrys(c echo.Context) error
	CreateCountry(c echo.Context) error
	UpdateCountry(c echo.Context) error
	DeleteCountry(c echo.Context) error
}

type countryHandler struct {
	usecase.CountryUseCase
}

func NewCountryHandler(u usecase.CountryUseCase) CountryHandler {
	return &countryHandler{u}
}

func (h *countryHandler) GetCountrys(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.CountryUseCase.GetCountrys(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Countrys does not exist.")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *countryHandler) CreateCountry(c echo.Context) error {
	object := &model.Countrys{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Countryname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.CountryUseCase.CreateCountry(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *countryHandler) UpdateCountry(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}
	object := &model.Countrys{
		Countryid: id,
	}
	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Countryid == 0 || object.Countryname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err = h.CountryUseCase.UpdateCountry(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, object)
	}

	return c.JSON(http.StatusOK, object)
}

func (h *countryHandler) DeleteCountry(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.CountryUseCase.DeleteCountry(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}
	return c.JSON(http.StatusOK, true)

}
