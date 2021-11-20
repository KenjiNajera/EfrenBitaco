package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	// GetRoleByID(c echo.Context) error
	GetRoles(c echo.Context) error
	CreateRole(c echo.Context) error
	UpdateRole(c echo.Context) error
	DeleteRole(c echo.Context) error
}

type roleHandler struct {
	RoleUseCase usecase.RoleUseCase
}

func NewRoleHandler(u usecase.RoleUseCase) RoleHandler {
	return &roleHandler{u}
}

func (h *roleHandler) GetRoles(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.RoleUseCase.GetRoles(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Roles does not exist.")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *roleHandler) CreateRole(c echo.Context) error {
	object := &model.Role{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Rolename == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err := h.RoleUseCase.CreateRole(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, object)
	}

	return c.JSON(http.StatusCreated, object)
}

func (h *roleHandler) UpdateRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Role ID must be int")
	}
	object := &model.Role{
		Roleid: id,
	}
	if err := c.Bind(object); err != nil {
		return err
	}
	if object.Roleid == 0 || object.Rolename == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err = h.RoleUseCase.UpdateRole(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Role can not Create.")
	}

	return c.JSON(http.StatusOK, object)
}

func (h *roleHandler) DeleteRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Role ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.RoleUseCase.DeleteRole(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Role can not Delete.")
	}

	return c.NoContent(http.StatusNoContent)
}
