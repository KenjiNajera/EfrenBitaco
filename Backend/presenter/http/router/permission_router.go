package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakePermissionHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("permission/permissions/", h.GetPermissions)
	r.POST("permission/create/", h.CreatePermission)
	r.PUT("permission/update/:id", h.UpdatePermission)
	r.DELETE("permission/delete/:id", h.DeletePermission)
}
