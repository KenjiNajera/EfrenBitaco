package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/grupogit/RMNGR002001/Backend/database"
	"github.com/grupogit/RMNGR002001/Backend/interactor"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/middleware"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/router"

	_ "github.com/grupogit/RMNGR002001/Backend/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Grupo GIT API
// @version 1.0
// @description This is a api rest full.
// @termsOfService http://swagger.io/terms/

// @contact.name Grupo GIT
// @contact.url http://www.grupogit.com/
// @contact.email contacto@grupogit.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2

func main() {
	db := database.InitDB()

	r := echo.New()

	defer func() {
		if err := db.Close(); err != nil {
			r.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()

	i := interactor.NewInteractor(db)

	h := i.NewAppHandler()

	//routes 33 funciones
	router.MakeBinnacleHandler(r, h)
	router.MakeSoftskillHandler(r, h)
	router.MakeHardskillHandler(r, h)
	router.MakeTypeRequestHandler(r, h)
	router.MakePermissionHandler(r, h)
	router.MakeProjectHandler(r, h)
	router.MakeModuleHandler(r, h)
	router.MakeAuthenticateHandler(r, h)
	router.MakeLanguageHandler(r, h)
	router.MakeUserHandler(r, h)
	router.MakeClientHandler(r, h)
	router.MakeCountryHandler(r, h)
	router.MakeRequestHandler(r, h)
	router.MakeRoleHandler(r, h)

	middleware.NewMiddleware(r)

	// Swagger
	r.GET("/*", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	// r.Static("/", "helper"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, r)
}
