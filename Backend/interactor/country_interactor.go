package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewCountryRepository() repository.CountryRepository {
	return datastore.NewCountryRepository(i.db)
}

// func (i *interactor) NewClientService() services.ClientService {
// 	return services.NewClientService(i.NewClientRepository())
// }

func (i *interactor) NewCountryUseCase() usecase.CountryUseCase {
	return usecase.NewCountryUseCase(i.NewCountryRepository())
}

func (i *interactor) NewCountryHandler() handler.CountryHandler {
	return handler.NewCountryHandler(i.NewCountryUseCase())
}
