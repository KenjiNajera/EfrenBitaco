package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewRoleRepository() repository.RoleRepository {
	return datastore.NewRoleRepository(i.db)
}

// func (i *interactor) NewRoleService() services.RoleService {
// 	return services.NewRoleService(i.NewRoleRepository())
// }

func (i *interactor) NewRoleUseCase() usecase.RoleUseCase {
	return usecase.NewRoleUseCase(i.NewRoleRepository())
}

func (i *interactor) NewRoleHandler() handler.RoleHandler {
	return handler.NewRoleHandler(i.NewRoleUseCase())
}
