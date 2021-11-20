package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type typerequestRepository struct {
	db *gorm.DB
}

func NewTyperequestRepository(db *gorm.DB) repository.TyperequestRepository {
	return &typerequestRepository{db: db}
}

func (r *typerequestRepository) GetTyperequests(ctx context.Context) ([]*model.Typerequest, error) {
	objects := []*model.Typerequest{}
	if err := r.db.Raw("SELECT * FROM typerequests").Scan(&objects).Error; err != nil {
		return objects, err
	}
	return objects, nil
}

func (r *typerequestRepository) CreateTyperequest(ctx context.Context, object *model.Typerequest) (int, error) {
	if err := r.db.Create(&object).Error; err != nil {
		return 0, err
	}
	return object.Typerequestid, nil
}

func (r *typerequestRepository) UpdateTyperequest(ctx context.Context, object *model.Typerequest) (bool, error) {
	sql := "UPDATE typerequests SET typerequestname = $1 WHERE typerequestid = $2"
	if err := r.db.Exec(sql, object.Typerequestname, object.Typerequestid).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *typerequestRepository) DeleteTyperequest(ctx context.Context, id int) error {
	sql := "DELETE FROM typerequests WHERE typerequestid = $1"
	err := r.db.Exec(sql, id).Error
	return err
}
