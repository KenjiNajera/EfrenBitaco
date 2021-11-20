package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type PermissionRepository interface {
	GetPermissions(ctx context.Context) ([]*model.Permission, error)
	CreatePermission(ctx context.Context, object *model.Permission) (int, error)
	UpdatePermission(ctx context.Context, object *model.Permission) (bool, error)
	DeletePermission(ctx context.Context, id int) error
}
