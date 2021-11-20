package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeAuthenticateHandler(r *echo.Echo, h handler.AppHandler) {
	r.POST("auth/authenticate/", h.Authenticate)
	r.POST("auth/recovery+password/", h.RecoveryPassword)
	r.POST("auth/change+password/", h.ChangePassword)
	r.POST("auth/verify+account/", h.VerifyAccount)
	r.POST("auth/authenticatemovil/", h.AuthenticateMovil)
}
