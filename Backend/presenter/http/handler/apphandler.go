package handler

type AppHandler interface {
	AuthenticateHandler
	BinnacleHandler
	ClientHandler
	CountryHandler
	HardskillHandler
	LanguageHandler
	ModuleHandler
	PermissionHandler
	ProjectHandler
	RequestHandler
	RoleHandler
	SoftskillHandler
	TyperequestHandler
	UserHandler
}
