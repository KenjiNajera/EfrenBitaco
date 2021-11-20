package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type RoleRepository interface {
	GetRoles(ctx context.Context) ([]*model.Role, error)
	CreateRole(ctx context.Context, object *model.Role) (int, error)
	UpdateRole(ctx context.Context, object *model.Role) (bool, error)
	DeleteRole(ctx context.Context, id int) error
}
