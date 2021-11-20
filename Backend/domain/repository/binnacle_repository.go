package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type BinnacleRepository interface {
	CreateBinnacleByID(ctx context.Context, CreateBinnacle *model.CreateBinnacle) error
	FinishBinnacleByID(ctx context.Context, id int) error
	GetBinnacleById(ctx context.Context, id int) (*model.BinnacleView, error)
	GetBinnacleByIdMovil(ctx context.Context, id int) (*model.BinnacleViewUser, error)
	CheckIn(ctx context.Context, checkin *model.Checktime) error
	CheckOut(ctx context.Context, departuretime *model.Departuretime) error
	MealDay(ctx context.Context, mealtime *model.Mealtime) error
	UpdateDay(ctx context.Context, data *model.UpdateDay) error
}
