package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type ProjectHandler interface {
	GetProjects(c echo.Context) error
	GetProjectByID(c echo.Context) error
	CreateProject(c echo.Context) error
	GetResources(c echo.Context) error
	GetLeaders(c echo.Context) error
	AllocateResource(c echo.Context) error
	AllocateLeader(c echo.Context) error
	DeleteResource(c echo.Context) error
	FinishProject(c echo.Context) error
	UpdateProject(c echo.Context) error
}

type projectHandler struct {
	usecase.ProjectUseCase
}

func NewProjectHandler(u usecase.ProjectUseCase) ProjectHandler {
	return &projectHandler{u}
}

func (h *projectHandler) GetProjects(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.GetProjects(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *projectHandler) GetProjectByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Project ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.GetProjectByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *projectHandler) CreateProject(c echo.Context) error {
	object := &model.Project{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Projectname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.CreateProject(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *projectHandler) GetResources(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.GetResources(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *projectHandler) GetLeaders(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.GetLeaders(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, result)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *projectHandler) AllocateResource(c echo.Context) error {
	object := &model.Member{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Userid == 0 || object.Projectid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.AllocateResource(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *projectHandler) AllocateLeader(c echo.Context) error {
	object := &model.Member{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Userid == 0 || object.Projectid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.AllocateLeader(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *projectHandler) DeleteResource(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Resource ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.ProjectUseCase.DeleteResource(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *projectHandler) FinishProject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Project ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = h.ProjectUseCase.FinishProject(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusCreated, true)
}

func (h *projectHandler) UpdateProject(c echo.Context) error {

	return c.JSON(http.StatusOK, true)
}
