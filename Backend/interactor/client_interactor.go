package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewClientRepository() repository.ClientRepository {
	return datastore.NewClientRepository(i.db)
}

// func (i *interactor) NewClientService() services.ClientService {
// 	return services.NewClientService(i.NewClientRepository())
// }

func (i *interactor) NewClientUseCase() usecase.ClientUseCase {
	return usecase.NewClientUseCase(i.NewClientRepository())
}

func (i *interactor) NewClientHandler() handler.ClientHandler {
	return handler.NewClientHandler(i.NewClientUseCase())
}
