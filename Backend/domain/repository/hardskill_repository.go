package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type HardskillRepository interface {
	GetHardskills(ctx context.Context) ([]*model.Hardskill, error)
	CreateHardskill(ctx context.Context, object *model.Hardskill) (int, error)
	UpdateHardskill(ctx context.Context, object *model.Hardskill) (bool, error)
	DeleteHardskill(ctx context.Context, id int) error
}
