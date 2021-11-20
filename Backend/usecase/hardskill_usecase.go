package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type HardskillUseCase interface {
	GetHardskills(ctx context.Context) ([]*model.Hardskill, error)
	CreateHardskill(ctx context.Context, object *model.Hardskill) (int, error)
	UpdateHardskill(ctx context.Context, object *model.Hardskill) (bool, error)
	DeleteHardskill(ctx context.Context, id int) error
}

type hardskillUseCase struct {
	repository.HardskillRepository
}

func NewHardskillUseCase(r repository.HardskillRepository) HardskillUseCase {
	return &hardskillUseCase{r}
}

func (u *hardskillUseCase) GetHardskills(ctx context.Context) ([]*model.Hardskill, error) {
	return u.HardskillRepository.GetHardskills(ctx)
}

func (u *hardskillUseCase) CreateHardskill(ctx context.Context, object *model.Hardskill) (int, error) {
	return u.HardskillRepository.CreateHardskill(ctx, object)
}

func (u *hardskillUseCase) UpdateHardskill(ctx context.Context, object *model.Hardskill) (bool, error) {
	return u.HardskillRepository.UpdateHardskill(ctx, object)
}

func (u *hardskillUseCase) DeleteHardskill(ctx context.Context, id int) error {
	return u.HardskillRepository.DeleteHardskill(ctx, id)
}
