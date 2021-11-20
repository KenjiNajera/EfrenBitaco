package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewAuthenticateRepository() repository.AuthenticateRepository {
	return datastore.NewAuthenticateRepository(i.db)
}

// func (i *interactor) NewAuthenticateService() services.AuthenticateService {
// 	return services.NewAuthenticateService(i.NewAuthenticateRepository())
// }

func (i *interactor) NewAuthenticateUseCase() usecase.AuthenticateUseCase {
	return usecase.NewAuthenticateUseCase(i.NewAuthenticateRepository())
}

func (i *interactor) NewAuthenticateHandler() handler.AuthenticateHandler {
	return handler.NewAuthenticateHandler(i.NewAuthenticateUseCase())
}
