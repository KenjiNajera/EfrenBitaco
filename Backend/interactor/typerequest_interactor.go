package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewTyperequestRepository() repository.TyperequestRepository {
	return datastore.NewTyperequestRepository(i.db)
}

// func (i *interactor) NewTyperequestService() services.TyperequestService {
// 	return services.NewTyperequestService(i.NewTyperequestRepository())
// }

func (i *interactor) NewTyperequestUseCase() usecase.TyperequestUseCase {
	return usecase.NewTyperequestUseCase(i.NewTyperequestRepository())
}

func (i *interactor) NewTyperequestHandler() handler.TyperequestHandler {
	return handler.NewTyperequestHandler(i.NewTyperequestUseCase())
}
