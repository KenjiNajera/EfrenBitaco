package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeBinnacleHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("binnacle/getbinnacle/:id", h.GetBinnacleById)
	r.GET("binnacle/getbinnacleuser/:id", h.GetBinnacleByIdMovil)
	r.POST("binnacle/create/", h.CreateBinnacleByID)
	r.POST("binnacle/checkin/", h.CheckIn)
	r.POST("binnacle/mealtime/", h.MealDay)
	r.POST("binnacle/checkout/", h.CheckOut)
}
