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

type HardskillHandler interface {
	GetHardskills(c echo.Context) error
	CreateHardskill(c echo.Context) error
	UpdateHardskill(c echo.Context) error
	DeleteHardskill(c echo.Context) error
}

type hardskillHandler struct {
	HardskillUseCase usecase.HardskillUseCase
}

func NewHardskillHandler(u usecase.HardskillUseCase) HardskillHandler {
	return &hardskillHandler{u}
}

func (h *hardskillHandler) GetHardskills(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.HardskillUseCase.GetHardskills(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Hardskills does not exist.")
	}

	return c.JSON(http.StatusOK, result)
}

func (h *hardskillHandler) CreateHardskill(c echo.Context) error {
	object := &model.Hardskill{}
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	str, errs := helper.Upload(file)
	if errs != nil {
		return err
	}
	object.Image = str

	data := c.FormValue("data")
	err = json.Unmarshal([]byte(data), object)
	if err != nil {
		fmt.Println(err)
	}

	if object.Hardskillname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.HardskillUseCase.CreateHardskill(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, object)
	}

	object.Hardskillid = result

	return c.JSON(http.StatusCreated, object)
}

func (h *hardskillHandler) UpdateHardskill(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}
	file, _ := c.FormFile("Image")

	str := c.FormValue("image")

	skill := c.FormValue("Hardskillname")
	object := &model.Hardskill{
		Hardskillid:   id,
		Image:         str,
		Hardskillname: skill,
	}

	if file != nil {
		object.Image, err = helper.Upload(file)
		if err != nil {
			return err
		}
	}

	if object.Hardskillid == 0 || object.Hardskillname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.HardskillUseCase.UpdateHardskill(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *hardskillHandler) DeleteHardskill(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.HardskillUseCase.DeleteHardskill(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
