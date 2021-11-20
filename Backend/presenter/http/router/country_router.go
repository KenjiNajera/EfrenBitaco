package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeCountryHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("country/getcountrys/", h.GetCountrys)
	r.POST("country/create/", h.CreateCountry)
	r.PUT("country/update/", h.UpdateCountry)
	r.DELETE("country/delete/:id", h.DeleteCountry)
}
