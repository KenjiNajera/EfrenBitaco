package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeHardskillHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("hardskill/hardskills/", h.GetHardskills)
	r.POST("hardskill/create/", h.CreateHardskill)
	r.PUT("hardskill/update/:id", h.UpdateHardskill)
	r.DELETE("hardskill/delete/:id", h.DeleteHardskill)
}
