package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewMiddleware(r *echo.Echo) {

	r.Use(
		middleware.Logger(),
		middleware.Recover(),
		// middleware.Timeout(60*time.Second),
		middleware.StaticWithConfig(middleware.StaticConfig{
			Root: "helper/images",
		}),
		middleware.Static(""),
		middleware.Static("helper/images"),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}),
	)
}
