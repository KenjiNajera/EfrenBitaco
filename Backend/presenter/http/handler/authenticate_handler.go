package handler

import (
	"context"
	"net/http"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
	"github.com/labstack/echo/v4"
)

type AuthenticateHandler interface {
	Authenticate(c echo.Context) error
	RecoveryPassword(c echo.Context) error
	ChangePassword(c echo.Context) error
	VerifyAccount(c echo.Context) error
	AuthenticateMovil(c echo.Context) error
}

type authenticateHandler struct {
	AuthenticateUseCase usecase.AuthenticateUseCase
}

func NewAuthenticateHandler(u usecase.AuthenticateUseCase) AuthenticateHandler {
	return &authenticateHandler{u}
}

// Authenticate godoc
// @Summary Authenticación de usuarios
// @Description get datos del usuario
// @Tags auth
// @Accept  json
// @Produce  json
// @Param objeto body model.Login true "Necesita un objeto"
// @Success 200 {object} model.Authenticate
// @Failure 400,404,500,default {string} string	"error"
// @Router /auth/authenticate/ [post]
func (h *authenticateHandler) Authenticate(c echo.Context) error {
	login := &model.Login{}

	if err := c.Bind(login); err != nil {
		return err
	}

	if login.Email == "" || login.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.AuthenticateUseCase.Authenticate(ctx, login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

// RecoveyPassword godoc
// @Summary Recuperacion de contraseñas
// @Description Envia un email al correo del usuario solicitante
// @Tags auth
// @Accept  json
// @Produce  json
// @Param objeto body model.RecoveryPassword true "Necesita un objeto"
// @Success 200 {boolean} bool "Retorna un true"
// @Failure 400,404,500,default {string} string	"error"
// @Router /auth/recovery+password/ [post]
func (h *authenticateHandler) RecoveryPassword(c echo.Context) error {
	email := &model.RecoveryPassword{}
	if err := c.Bind(email); err != nil {
		return err
	}
	if email.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.AuthenticateUseCase.RecoveryPassword(ctx, email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// ChangePassword godoc
// @Summary Cambia la contraseña del usuario
// @Description Recibe el id hasheado del solicitante de cambio de contraseña y valida si el usario ya ha pedido cambiar su contraseña o no.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param objeto body model.ChangePassword true "Necesita un objeto"
// @Success 200 {boolean} bool "Retorna un true"
// @Failure 400,404,500,default {string} string	"error"
// @Router /auth/change+password/ [post]
func (h *authenticateHandler) ChangePassword(c echo.Context) error {
	changepassword := &model.ChangePassword{}
	if err := c.Bind(changepassword); err != nil {
		return err
	}

	if changepassword.Hashuserid == "" || changepassword.Newpassword == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.AuthenticateUseCase.ChangePassword(ctx, changepassword)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// AuthenticateMovil godoc
// @Summary Authenticación de usuarios en app movil
// @Description get datos del usuario en la aplicación movil
// @Tags auth
// @Accept  json
// @Produce  json
// @Param objeto body model.LoginMovil true "Necesita un objeto"
// @Success 200 {object} model.AuthenticateMovil
// @Failure 400,404,500,default {string} string	"error"
// @Router /auth/authenticatemovil/ [post]
func (h *authenticateHandler) AuthenticateMovil(c echo.Context) error {
	loginmovil := &model.LoginMovil{}

	if err := c.Bind(loginmovil); err != nil {
		return err
	}

	if loginmovil.Email == "" || loginmovil.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.AuthenticateUseCase.AuthenticateMovil(ctx, loginmovil)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// VerifyAccount godoc
// @Summary Verificación de Usaurios nuevos
// @Description Recibe el id hasheado del solicitante de verificación de cuenta y valida si el usario ya esta verificado o no.s
// @Tags auth
// @Accept  json
// @Produce  json
// @Param objeto body model.ChangePassword true "Necesita un objeto"
// @Success 200 {boolean} bool "Retorna un true"
// @Failure 400,404,500,default {string} string	"error"
// @Router /auth/verify+account/ [post]
func (h *authenticateHandler) VerifyAccount(c echo.Context) error {
	changepassword := &model.ChangePassword{}
	if err := c.Bind(changepassword); err != nil {
		return err
	}

	if changepassword.Hashuserid == "" || changepassword.Newpassword == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Syntax error of sent json")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := h.AuthenticateUseCase.VerifyAccount(ctx, changepassword)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
