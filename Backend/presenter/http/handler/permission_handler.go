package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type PermissionHandler interface {
	GetPermissions(c echo.Context) error
	CreatePermission(c echo.Context) error
	UpdatePermission(c echo.Context) error
	DeletePermission(c echo.Context) error
}

type permissionHandler struct {
	PermissionUseCase usecase.PermissionUseCase
}

func NewPermissionHandler(u usecase.PermissionUseCase) PermissionHandler {
	return &permissionHandler{u}
}

func (h *permissionHandler) GetPermissions(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.PermissionUseCase.GetPermissions(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Permissions does not exist.")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *permissionHandler) CreatePermission(c echo.Context) error {
	object := &model.Permission{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Permissionname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.PermissionUseCase.CreatePermission(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *permissionHandler) UpdatePermission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}
	object := &model.Permission{
		Permissionid: id,
	}
	if err := c.Bind(object); err != nil {
		return err
	}
	if object.Permissionid == 0 || object.Permissionname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.PermissionUseCase.UpdatePermission(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *permissionHandler) DeletePermission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.PermissionUseCase.DeletePermission(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
