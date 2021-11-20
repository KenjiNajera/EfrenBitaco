package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewPermissionRepository() repository.PermissionRepository {
	return datastore.NewPermissionRepository(i.db)
}

// func (i *interactor) NewPermissionService() services.PermissionService {
// 	return services.NewPermissionService(i.NewPermissionRepository())
// }

func (i *interactor) NewPermissionUseCase() usecase.PermissionUseCase {
	return usecase.NewPermissionUseCase(i.NewPermissionRepository())
}

func (i *interactor) NewPermissionHandler() handler.PermissionHandler {
	return handler.NewPermissionHandler(i.NewPermissionUseCase())
}
