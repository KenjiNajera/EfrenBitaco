package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewRequestRepository() repository.RequestRepository {
	return datastore.NewRequestRepository(i.db)
}

// func (i *interactor) NewTyperequestService() services.TyperequestService {
// 	return services.NewTyperequestService(i.NewTyperequestRepository())
// }

func (i *interactor) NewRequestUseCase() usecase.RequestUseCase {
	return usecase.NewRequestUseCase(i.NewRequestRepository())
}

func (i *interactor) NewRequestHandler() handler.RequestHandler {
	return handler.NewRequestHandler(i.NewRequestUseCase())
}
