package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type SoftskillRepository interface {
	GetSoftskills(ctx context.Context) ([]*model.Softskill, error)
	CreateSoftskill(ctx context.Context, object *model.Softskill) (int, error)
	UpdateSoftskill(ctx context.Context, object *model.Softskill) (bool, error)
	DeleteSoftskill(ctx context.Context, id int) error
}
