package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewModuleRepository() repository.ModuleRepository {
	return datastore.NewModuleRepository(i.db)
}

// func (i *interactor) NewModuleService() services.ModuleService {
// 	return services.NewModuleService(i.NewModuleRepository())
// }

func (i *interactor) NewModuleUseCase() usecase.ModuleUseCase {
	return usecase.NewModuleUseCase(i.NewModuleRepository())
}

func (i *interactor) NewModuleHandler() handler.ModuleHandler {
	return handler.NewModuleHandler(i.NewModuleUseCase())
}
