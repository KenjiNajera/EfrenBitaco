package router

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/labstack/echo/v4"
)

func MakeLanguageHandler(r *echo.Echo, h handler.AppHandler) {
	r.GET("language/languages/", h.GetLanguages)
	r.POST("language/create/", h.CreateLanguage)
	r.PUT("language/update/:id", h.UpdateLanguage)
	r.DELETE("language/delete/:id", h.DeleteLanguage)
}
