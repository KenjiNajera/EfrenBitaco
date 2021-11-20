package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) GetPermissions(ctx context.Context) ([]*model.Permission, error) {
	objects := []*model.Permission{}
	if err := r.db.Raw("SELECT * FROM permissions").Scan(&objects).Error; err != nil {
		return objects, err
	}
	return objects, nil
}

func (r *permissionRepository) CreatePermission(ctx context.Context, object *model.Permission) (int, error) {
	if err := r.db.Create(&object).Error; err != nil {
		return 0, err
	}
	return object.Permissionid, nil
}

func (r *permissionRepository) UpdatePermission(ctx context.Context, object *model.Permission) (bool, error) {
	sql := "UPDATE permissions SET permissionname = $1 WHERE permissionid = $2"
	if err := r.db.Exec(sql, object.Permissionname, object.Permissionid).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *permissionRepository) DeletePermission(ctx context.Context, id int) error {
	sql := "DELETE FROM permissions WHERE ermissionid = $1"
	err := r.db.Exec(sql, id).Error
	return err
}
