package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type ModuleRepository interface {
	GetModules(ctx context.Context) ([]*model.Module, error)
	CreateModule(ctx context.Context, object *model.Module) (int, error)
	UpdateModule(ctx context.Context, object *model.Module) (bool, error)
	DeleteModule(ctx context.Context, id int) error
}
