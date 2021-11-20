package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type RequestHandler interface {
	GetRequests(c echo.Context) error
	CreateRequest(c echo.Context) error
	ResponseRequest(c echo.Context) error
	GetRequestsUser(c echo.Context) error
}

type requestHandler struct {
	usecase.RequestUseCase
}

func NewRequestHandler(u usecase.RequestUseCase) RequestHandler {
	return &requestHandler{u}
}

func (h *requestHandler) GetRequests(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.RequestUseCase.GetRequests(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *requestHandler) CreateRequest(c echo.Context) error {
	object := &model.Request{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Description == "" || object.Userid == 0 || object.Typerequestid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.RequestUseCase.CreateRequest(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *requestHandler) ResponseRequest(c echo.Context) error {
	object := &model.Request{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Details == "" || object.Requestid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.RequestUseCase.ResponseRequest(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *requestHandler) GetRequestsUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Request ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.RequestUseCase.GetRequestsUser(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}
