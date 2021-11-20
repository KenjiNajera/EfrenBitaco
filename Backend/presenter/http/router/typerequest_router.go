package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeTypeRequestHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("typerequest/gettyperequests/", h.GetTyperequests)
	r.POST("typerequest/create/", h.CreateTyperequest)
	r.PUT("typerequest/update/:id", h.UpdateTyperequest)
	r.DELETE("typerequest/delete/:id", h.DeleteTyperequest)
}
