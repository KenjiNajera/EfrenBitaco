package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) repository.CountryRepository {
	return &countryRepository{db}
}

func (r *countryRepository) GetCountrys(ctx context.Context) ([]*model.Countrys, error) {
	objects := []*model.Countrys{}
	sqlGetCountrys := `SELECT * FROM countrys`
	if err := r.db.Raw(sqlGetCountrys).Scan(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func (r *countryRepository) CreateCountry(ctx context.Context, object *model.Countrys) (int, error) {
	if err := r.db.Create(&object).Error; err != nil {
		return 0, err
	}
	return object.Countryid, nil
}

func (r *countryRepository) UpdateCountry(ctx context.Context, object *model.Countrys) (bool, error) {
	sql := "UPDATE countrys SET countryname = $1 WHERE countryid = $2"
	if err := r.db.Exec(sql, object.Countryname, object.Countryid).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *countryRepository) DeleteCountry(ctx context.Context, id int) error {
	sql := "DELETE FROM countrys WHERE countryid = $1"
	if err := r.db.Exec(sql, id).Error; err != nil {
		return err
	}
	return nil
}
