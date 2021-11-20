package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"

	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type SoftskillUseCase interface {
	GetSoftskills(ctx context.Context) ([]*model.Softskill, error)
	CreateSoftskill(ctx context.Context, skill *model.Softskill) (int, error)
	UpdateSoftskill(ctx context.Context, skill *model.Softskill) (bool, error)
	DeleteSoftskill(ctx context.Context, id int) error
}

type softskillUseCase struct {
	repository.SoftskillRepository
}

func NewSoftskillUseCase(r repository.SoftskillRepository) SoftskillUseCase {
	return &softskillUseCase{r}
}

func (u *softskillUseCase) GetSoftskills(ctx context.Context) ([]*model.Softskill, error) {
	return u.SoftskillRepository.GetSoftskills(ctx)
}

func (u *softskillUseCase) CreateSoftskill(ctx context.Context, user *model.Softskill) (int, error) {
	return u.SoftskillRepository.CreateSoftskill(ctx, user)
}

func (u *softskillUseCase) UpdateSoftskill(ctx context.Context, skill *model.Softskill) (bool, error) {
	return u.SoftskillRepository.UpdateSoftskill(ctx, skill)
}

func (u *softskillUseCase) DeleteSoftskill(ctx context.Context, id int) error {
	return u.SoftskillRepository.DeleteSoftskill(ctx, id)
}
