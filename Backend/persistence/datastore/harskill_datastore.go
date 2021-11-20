package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type hardskillRepository struct {
	db *gorm.DB
}

func NewHardskillRepository(db *gorm.DB) repository.HardskillRepository {
	return &hardskillRepository{db: db}
}

func (r *hardskillRepository) GetHardskills(ctx context.Context) ([]*model.Hardskill, error) {
	objects := []*model.Hardskill{}
	if err := r.db.Raw("SELECT * FROM hardskills").Scan(&objects).Error; err != nil {
		return objects, err
	}
	return objects, nil
}

func (r *hardskillRepository) CreateHardskill(ctx context.Context, object *model.Hardskill) (int, error) {
	if err := r.db.Create(&object).Error; err != nil {
		return 0, err
	}
	return object.Hardskillid, nil
}

func (r *hardskillRepository) UpdateHardskill(ctx context.Context, object *model.Hardskill) (bool, error) {
	sql := "UPDATE hardskills SET hardskillname = $1, image = $2 WHERE hardskillid = $3"
	if err := r.db.Exec(sql, object.Hardskillname, object.Image, object.Hardskillid).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *hardskillRepository) DeleteHardskill(ctx context.Context, id int) error {
	sql := "DELETE FROM hardskills WHERE hardskillid = $1"
	err := r.db.Exec(sql, id).Error
	return err
}
