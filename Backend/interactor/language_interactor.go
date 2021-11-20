package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewLanguageRepository() repository.LanguageRepository {
	return datastore.NewLanguageRepository(i.db)
}

// func (i *interactor) NewLanguageService() services.LanguageService {
// 	return services.NewLanguageService(i.NewLanguageRepository())
// }

func (i *interactor) NewLanguageUseCase() usecase.LanguageUseCase {
	return usecase.NewLanguageUseCase(i.NewLanguageRepository())
}

func (i *interactor) NewLanguageHandler() handler.LanguageHandler {
	return handler.NewLanguageHandler(i.NewLanguageUseCase())
}
