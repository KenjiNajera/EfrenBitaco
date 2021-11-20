package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeRequestHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("request/getrequests/", h.GetRequests)
	r.POST("request/create/", h.CreateRequest)
	r.POST("request/response/", h.ResponseRequest)
	r.GET("request/getrequestuser/:id", h.GetRequestsUser)
}
