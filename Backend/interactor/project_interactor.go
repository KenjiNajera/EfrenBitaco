package interactor

import (
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/persistence/datastore"
	"github.com/grupogit/RMNGR002001/Backend/presenter/http/handler"
	"github.com/grupogit/RMNGR002001/Backend/usecase"
)

func (i *interactor) NewProjectRepository() repository.ProjectRepository {
	return datastore.NewProjectRepository(i.db)
}

// func (i *interactor) NewHardskillService() services.HardskillService {
// 	return services.NewHardskillService(i.NewHardskillRepository())
// }

func (i *interactor) NewProjectUseCase() usecase.ProjectUseCase {
	return usecase.NewProjectUseCase(i.NewProjectRepository())
}

func (i *interactor) NewProjectHandler() handler.ProjectHandler {
	return handler.NewProjectHandler(i.NewProjectUseCase())
}
