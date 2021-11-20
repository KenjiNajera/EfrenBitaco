package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeProjectHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("project/getprojects/", h.GetProjects)
	r.GET("project/getproject/:id", h.GetProjectByID)
	r.POST("project/create/", h.CreateProject)
	r.GET("project/getresources/", h.GetResources)
	r.GET("project/getleaders/", h.GetLeaders)
	r.POST("project/allocateresource/", h.AllocateResource)
	r.POST("project/allocateleader/", h.AllocateLeader)
	r.DELETE("project/deleteresouce/:id", h.DeleteResource)
	r.DELETE("project/finishproject/:id", h.FinishProject)
	r.PUT("project/update/", h.UpdateProject)
}
