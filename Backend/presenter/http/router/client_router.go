package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeClientHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("client/clients/", h.GetClients)
	r.GET("client/client/:id", h.GetClientByID)
	r.POST("client/create/", h.CreateClient)
	r.PUT("client/update/", h.UpdateClient)
	r.DELETE("client/delete/:id", h.DeleteClient)
	r.DELETE("client/reactivate/:id", h.Reactivate)
}
