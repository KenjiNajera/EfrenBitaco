package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type RoleUseCase interface {
	GetRoles(ctx context.Context) ([]*model.Role, error)
	CreateRole(ctx context.Context, object *model.Role) (int, error)
	UpdateRole(ctx context.Context, object *model.Role) (bool, error)
	DeleteRole(ctx context.Context, id int) error
}

type roleUseCase struct {
	repository.RoleRepository
}

func NewRoleUseCase(r repository.RoleRepository) RoleUseCase {
	return &roleUseCase{r}
}

func (u *roleUseCase) GetRoles(ctx context.Context) ([]*model.Role, error) {
	return u.RoleRepository.GetRoles(ctx)
}

func (u *roleUseCase) CreateRole(ctx context.Context, object *model.Role) (int, error) {
	return u.RoleRepository.CreateRole(ctx, object)
}

func (u *roleUseCase) UpdateRole(ctx context.Context, object *model.Role) (bool, error) {
	return u.RoleRepository.UpdateRole(ctx, object)
}

func (u *roleUseCase) DeleteRole(ctx context.Context, id int) error {
	return u.RoleRepository.DeleteRole(ctx, id)
}
