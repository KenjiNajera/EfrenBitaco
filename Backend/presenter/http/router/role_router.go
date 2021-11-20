package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeRoleHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("role/getroles/", h.GetRoles)
	r.POST("role/create/", h.CreateRole)
	r.PUT("role/update/", h.UpdateRole)
	r.DELETE("role/delete/:id", h.DeleteRole)
}
