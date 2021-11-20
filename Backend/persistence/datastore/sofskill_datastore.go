package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type softskillRepository struct {
	db *gorm.DB
}

func NewSoftskillRepository(db *gorm.DB) repository.SoftskillRepository {
	return &softskillRepository{db: db}
}

func (r *softskillRepository) GetSoftskills(ctx context.Context) ([]*model.Softskill, error) {
	objects := []*model.Softskill{}
	if err := r.db.Raw("SELECT * FROM softskills").Scan(&objects).Error; err != nil {
		return objects, err
	}
	return objects, nil
}

func (r *softskillRepository) CreateSoftskill(ctx context.Context, object *model.Softskill) (int, error) {
	if err := r.db.Create(&object).Error; err != nil {
		return 0, err
	}
	return object.Softskillid, nil
}

func (r *softskillRepository) UpdateSoftskill(ctx context.Context, object *model.Softskill) (bool, error) {
	sql := "UPDATE softskills SET softskillname = $1 WHERE softskillid = $2"
	if err := r.db.Exec(sql, object.Softskillname, object.Softskillid).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *softskillRepository) DeleteSoftskill(ctx context.Context, id int) error {
	sql := "DELETE FROM softskills WHERE softskillid = $1"
	err := r.db.Exec(sql, id).Error
	return err
}
