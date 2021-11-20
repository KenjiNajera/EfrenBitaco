package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"

	"github.com/labstack/echo/v4"
)

func MakeSoftskillHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("softskill/getsoftskills/", h.GetSoftskills)
	r.POST("softskill/create/", h.CreateSoftskill)
	r.PUT("softskill/update/:id", h.UpdateSoftskill)
	r.DELETE("softskill/delete/:id", h.DeleteSoftskill)
}
