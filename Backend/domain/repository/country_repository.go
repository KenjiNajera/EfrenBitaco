package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type CountryRepository interface {
	GetCountrys(ctx context.Context) ([]*model.Countrys, error)
	CreateCountry(ctx context.Context, object *model.Countrys) (int, error)
	UpdateCountry(ctx context.Context, object *model.Countrys) (bool, error)
	DeleteCountry(ctx context.Context, id int) error
}
