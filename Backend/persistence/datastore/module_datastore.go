package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type moduleRepository struct {
	db *gorm.DB
}

func NewModuleRepository(db *gorm.DB) repository.ModuleRepository {
	return &moduleRepository{db: db}
}

func (r *moduleRepository) GetModules(ctx context.Context) ([]*model.Module, error) {
	objects := []*model.Module{}
	if err := r.db.Raw("SELECT * FROM modules").Scan(&objects).Error; err != nil {
		return objects, err
	}
	return objects, nil
}

func (r *moduleRepository) CreateModule(ctx context.Context, object *model.Module) (int, error) {
	if err := r.db.Create(&object).Error; err != nil {
		return 0, err
	}
	return object.Moduleid, nil
}

func (r *moduleRepository) UpdateModule(ctx context.Context, object *model.Module) (bool, error) {
	sql := "UPDATE modules SET modulename = $1 WHERE moduleid = $2"
	if err := r.db.Exec(sql, object.Modulename, object.Moduleid).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *moduleRepository) DeleteModule(ctx context.Context, id int) error {
	sql := "DELETE FROM modules WHERE moduleid = $1"
	err := r.db.Exec(sql, id).Error
	return err
}
