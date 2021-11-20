package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeModuleHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("module/modules/", h.GetModules)
	r.POST("module/create/", h.CreateModule)
	r.PUT("module/update/:id", h.UpdateModule)
	r.DELETE("module/delete/:id", h.DeleteModule)
}
