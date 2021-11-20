package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type LanguageRepository interface {
	GetLanguages(ctx context.Context) ([]*model.Language, error)
	CreateLanguage(ctx context.Context, object *model.Language) (int, error)
	UpdateLanguage(ctx context.Context, object *model.Language) (bool, error)
	DeleteLanguage(ctx context.Context, id int) error
}
