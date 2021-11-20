package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewHardskillRepository() repository.HardskillRepository {
	return datastore.NewHardskillRepository(i.db)
}

// func (i *interactor) NewHardskillService() services.HardskillService {
// 	return services.NewHardskillService(i.NewHardskillRepository())
// }

func (i *interactor) NewHardskillUseCase() usecase.HardskillUseCase {
	return usecase.NewHardskillUseCase(i.NewHardskillRepository())
}

func (i *interactor) NewHardskillHandler() handler.HardskillHandler {
	return handler.NewHardskillHandler(i.NewHardskillUseCase())
}
