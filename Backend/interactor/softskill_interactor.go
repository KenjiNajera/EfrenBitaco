package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewSoftskillRepository() repository.SoftskillRepository {
	return datastore.NewSoftskillRepository(i.db)
}

// func (i *interactor) NewSoftskillService() services.SoftskillService {
// 	return services.NewsoftskillService(i.NewSoftskillRepository())
// }

func (i *interactor) NewSoftskillUseCase() usecase.SoftskillUseCase {
	return usecase.NewSoftskillUseCase(i.NewSoftskillRepository())
}

func (i *interactor) NewSoftskillHandler() handler.SoftskillHandler {
	return handler.NewSoftskillHandler(i.NewSoftskillUseCase())
}
