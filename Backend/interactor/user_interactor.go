package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(i.db)
}

// func (i *interactor) NewCertificateService() services.CertificateService {
// 	return services.NewCertificateService(i.NewUserRepository())
// }

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}
