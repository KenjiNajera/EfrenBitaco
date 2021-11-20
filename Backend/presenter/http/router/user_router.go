package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeUserHandler(r *echo.Echo, h handler.AppHandler) {
	r.POST("user/register/", h.CreateUser)
	r.PUT("user/update/", h.UpdateUser)
	r.GET("user/activate/:id", h.ActivateUser)
	r.GET("user/desactivate/:id", h.DesactivateUser)
	r.GET("user/getusers/", h.GetUsers)
	r.GET("user/getmydatamovil/:id", h.GetMyDataMovil)
}
