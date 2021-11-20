package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewBinnacleRepository() repository.BinnacleRepository {
	return datastore.NewBinnacleRepository(i.db)
}

// func (i *interactor) NewAuthenticateService() services.AuthenticateService {
// 	return services.NewAuthenticateService(i.NewBinnacleRepository())
// }

func (i *interactor) NewBinnacleUseCase() usecase.BinnacleUseCase {
	return usecase.NewBinnacleUseCase(i.NewBinnacleRepository())
}

func (i *interactor) NewBinnacleHandler() handler.BinnacleHandler {
	return handler.NewBinnacleHandler(i.NewBinnacleUseCase())
}
