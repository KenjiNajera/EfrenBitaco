package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type languageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) repository.LanguageRepository {
	return &languageRepository{db: db}
}

func (r *languageRepository) GetLanguages(ctx context.Context) ([]*model.Language, error) {
	objects := []*model.Language{}
	if err := r.db.Raw("SELECT * FROM languages").Scan(&objects).Error; err != nil {
		return objects, err
	}
	return objects, nil
}

func (r *languageRepository) CreateLanguage(ctx context.Context, object *model.Language) (int, error) {
	if err := r.db.Create(&object).Error; err != nil {
		return 0, err
	}
	return object.Languageid, nil
}

func (r *languageRepository) UpdateLanguage(ctx context.Context, object *model.Language) (bool, error) {
	sql := "UPDATE languages SET languagename = $1 WHERE languageid = $2"
	if err := r.db.Exec(sql, object.Languagename, object.Languageid).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *languageRepository) DeleteLanguage(ctx context.Context, id int) error {
	sql := "DELETE FROM languages WHERE languageid = $1"
	err := r.db.Exec(sql, id).Error
	return err
}
