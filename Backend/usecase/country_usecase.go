package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type CountryUseCase interface {
	GetCountrys(ctx context.Context) ([]*model.Countrys, error)
	CreateCountry(ctx context.Context, object *model.Countrys) (int, error)
	UpdateCountry(ctx context.Context, object *model.Countrys) (bool, error)
	DeleteCountry(ctx context.Context, id int) error
}

type countryUseCase struct {
	repository.CountryRepository
}

func NewCountryUseCase(r repository.CountryRepository) CountryUseCase {
	return &countryUseCase{r}
}

func (u *countryUseCase) GetCountrys(ctx context.Context) ([]*model.Countrys, error) {
	return u.CountryRepository.GetCountrys(ctx)
}

func (u *countryUseCase) CreateCountry(ctx context.Context, object *model.Countrys) (int, error) {
	return u.CountryRepository.CreateCountry(ctx, object)
}

func (u *countryUseCase) UpdateCountry(ctx context.Context, object *model.Countrys) (bool, error) {
	return u.CountryRepository.UpdateCountry(ctx, object)
}

func (u *countryUseCase) DeleteCountry(ctx context.Context, id int) error {
	return u.CountryRepository.DeleteCountry(ctx, id)
}
