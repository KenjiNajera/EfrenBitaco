package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type LanguageHandler interface {
	GetLanguages(c echo.Context) error
	CreateLanguage(c echo.Context) error
	UpdateLanguage(c echo.Context) error
	DeleteLanguage(c echo.Context) error
}

type languageHandler struct {
	LanguageUseCase usecase.LanguageUseCase
}

func NewLanguageHandler(u usecase.LanguageUseCase) LanguageHandler {
	return &languageHandler{u}
}

// GetLanguages godoc
// @Summary Lista de los idiomas existentes.
// @Description Retorna un Array de Idiomas.
// @Tags language
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Language "Retorna un array de idiomas"
// @Failure 400,404,500,default {string} string	"error"
// @Router /language/languages/ [get]
func (h *languageHandler) GetLanguages(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	objects, err := h.LanguageUseCase.GetLanguages(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Languages does not exist.")
	}

	return c.JSON(http.StatusOK, objects)
}

// CreateLanguage godoc
// @Summary Agregar Idioma.
// @Description Agrega un nuevo idioma.
// @Tags language
// @Accept  json
// @Produce  json
// @Param objeto body model.Language true "Necesita un objeto"
// @Success 200 {integer} integer "Retorna el id del idioma creado"
// @Failure 400,404,500,default {string} string	"error"
// @Router /language/create/ [post]
func (h *languageHandler) CreateLanguage(c echo.Context) error {
	object := &model.Language{}

	if err := c.Bind(object); err != nil {
		return err
	}

	if object.Languagename == "" {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.LanguageUseCase.CreateLanguage(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusCreated, result)
}

// UpdateLanguage godoc
// @Summary Modificar datos del idioma.
// @Description Modifica el dato del idioma registrado.
// @Tags language
// @Accept  json
// @Produce  json
// @Param id path int true "Language ID"
// @Param objeto body model.Language true "Necesita un objeto"
// @Success 200 {boolean} boolean "Retorna un true"
// @Failure 400,404,500,default {string} string	"error"
// @Router /language/update/:id [put]
func (h *languageHandler) UpdateLanguage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}
	object := &model.Language{
		Languageid: id,
	}
	if err := c.Bind(object); err != nil {
		return err
	}
	if object.Languageid == 0 || object.Languagename == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.LanguageUseCase.UpdateLanguage(ctx, object)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)
}

// DeleteLanguage godoc
// @Summary Elimina idioma existente.
// @Description Elimina algun idioma Existente.
// @Tags language
// @Accept  json
// @Produce  json
// @Param id path int true "Language ID"
// @Success 200 {boolean} boolean "Retorna un true"
// @Failure 400,404,500,default {string} string	"error"
// @Router /language/delete/:id [delete]
func (h *languageHandler) DeleteLanguage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, false)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.LanguageUseCase.DeleteLanguage(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
