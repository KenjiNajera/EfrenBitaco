package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type ModuleUseCase interface {
	GetModules(ctx context.Context) ([]*model.Module, error)
	CreateModule(ctx context.Context, object *model.Module) (int, error)
	UpdateModule(ctx context.Context, object *model.Module) (bool, error)
	DeleteModule(ctx context.Context, id int) error
}

type moduleUseCase struct {
	repository.ModuleRepository
}

func NewModuleUseCase(r repository.ModuleRepository) ModuleUseCase {
	return &moduleUseCase{r}
}

func (u *moduleUseCase) GetModules(ctx context.Context) ([]*model.Module, error) {
	return u.ModuleRepository.GetModules(ctx)
}

func (u *moduleUseCase) CreateModule(ctx context.Context, object *model.Module) (int, error) {
	return u.ModuleRepository.CreateModule(ctx, object)
}

func (u *moduleUseCase) UpdateModule(ctx context.Context, object *model.Module) (bool, error) {
	return u.ModuleRepository.UpdateModule(ctx, object)
}

func (u *moduleUseCase) DeleteModule(ctx context.Context, id int) error {
	return u.ModuleRepository.DeleteModule(ctx, id)
}
