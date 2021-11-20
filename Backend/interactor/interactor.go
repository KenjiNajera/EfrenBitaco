package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/jinzhu/gorm"
)

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	db *gorm.DB
}

func NewInteractor(db *gorm.DB) Interactor {
	return &interactor{db}
}

type appHandler struct {
	handler.HardskillHandler
	handler.LanguageHandler
	handler.ModuleHandler
	handler.PermissionHandler
	handler.RoleHandler
	handler.SoftskillHandler
	handler.TyperequestHandler
	handler.AuthenticateHandler
	handler.ClientHandler
	handler.UserHandler
	handler.RequestHandler
	handler.ProjectHandler
	handler.CountryHandler
	handler.BinnacleHandler
	// embed all handler interfaces
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.AuthenticateHandler = i.NewAuthenticateHandler()
	appHandler.BinnacleHandler = i.NewBinnacleHandler()
	appHandler.ClientHandler = i.NewClientHandler()
	appHandler.CountryHandler = i.NewCountryHandler()
	appHandler.HardskillHandler = i.NewHardskillHandler()
	appHandler.LanguageHandler = i.NewLanguageHandler()
	appHandler.ModuleHandler = i.NewModuleHandler()
	appHandler.PermissionHandler = i.NewPermissionHandler()
	appHandler.ProjectHandler = i.NewProjectHandler()
	appHandler.RequestHandler = i.NewRequestHandler()
	appHandler.RoleHandler = i.NewRoleHandler()
	appHandler.SoftskillHandler = i.NewSoftskillHandler()
	appHandler.TyperequestHandler = i.NewTyperequestHandler()
	appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}
