package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type BinnacleUseCase interface {
	CreateBinnacleByID(ctx context.Context, CreateBinnacle *model.CreateBinnacle) error
	FinishBinnacleByID(ctx context.Context, id int) error
	GetBinnacleById(ctx context.Context, id int) (*model.BinnacleView, error)
	GetBinnacleByIdMovil(ctx context.Context, id int) (*model.BinnacleViewUser, error)
	CheckIn(ctx context.Context, checkin *model.Checktime) error
	CheckOut(ctx context.Context, departuretime *model.Departuretime) error
	MealDay(ctx context.Context, mealtime *model.Mealtime) error
	UpdateDay(ctx context.Context, data *model.UpdateDay) error
}

type binnacleUseCase struct {
	repository.BinnacleRepository
}

func NewBinnacleUseCase(r repository.BinnacleRepository) BinnacleUseCase {
	return &binnacleUseCase{r}
}

func (u *binnacleUseCase) CreateBinnacleByID(ctx context.Context, CreateBinnacle *model.CreateBinnacle) error {
	return u.BinnacleRepository.CreateBinnacleByID(ctx, CreateBinnacle)
}

func (u *binnacleUseCase) FinishBinnacleByID(ctx context.Context, id int) error {
	return u.BinnacleRepository.FinishBinnacleByID(ctx, id)
}

func (u *binnacleUseCase) GetBinnacleById(ctx context.Context, id int) (*model.BinnacleView, error) {
	return u.BinnacleRepository.GetBinnacleById(ctx, id)
}

func (u *binnacleUseCase) GetBinnacleByIdMovil(ctx context.Context, id int) (*model.BinnacleViewUser, error) {
	return u.BinnacleRepository.GetBinnacleByIdMovil(ctx, id)
}

func (u *binnacleUseCase) CheckIn(ctx context.Context, checkin *model.Checktime) error {
	return u.BinnacleRepository.CheckIn(ctx, checkin)
}

func (u *binnacleUseCase) CheckOut(ctx context.Context, departuretime *model.Departuretime) error {
	return u.BinnacleRepository.CheckOut(ctx, departuretime)
}

func (u *binnacleUseCase) MealDay(ctx context.Context, mealtime *model.Mealtime) error {
	return u.BinnacleRepository.MealDay(ctx, mealtime)
}

func (u *binnacleUseCase) UpdateDay(ctx context.Context, data *model.UpdateDay) error {
	return u.BinnacleRepository.UpdateDay(ctx, data)
}
