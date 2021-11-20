package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type PermissionUseCase interface {
	GetPermissions(ctx context.Context) ([]*model.Permission, error)
	CreatePermission(ctx context.Context, object *model.Permission) (int, error)
	UpdatePermission(ctx context.Context, object *model.Permission) (bool, error)
	DeletePermission(ctx context.Context, id int) error
}

type permissionUseCase struct {
	repository.PermissionRepository
}

func NewPermissionUseCase(r repository.PermissionRepository) PermissionUseCase {
	return &permissionUseCase{r}
}

func (u *permissionUseCase) GetPermissions(ctx context.Context) ([]*model.Permission, error) {
	return u.PermissionRepository.GetPermissions(ctx)
}

func (u *permissionUseCase) CreatePermission(ctx context.Context, object *model.Permission) (int, error) {
	return u.PermissionRepository.CreatePermission(ctx, object)
}

func (u *permissionUseCase) UpdatePermission(ctx context.Context, object *model.Permission) (bool, error) {
	return u.PermissionRepository.UpdatePermission(ctx, object)
}

func (u *permissionUseCase) DeletePermission(ctx context.Context, id int) error {
	return u.PermissionRepository.DeletePermission(ctx, id)
}
