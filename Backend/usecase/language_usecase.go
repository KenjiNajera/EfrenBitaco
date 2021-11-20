package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type LanguageUseCase interface {
	GetLanguages(ctx context.Context) ([]*model.Language, error)
	CreateLanguage(ctx context.Context, skill *model.Language) (int, error)
	UpdateLanguage(ctx context.Context, skill *model.Language) (bool, error)
	DeleteLanguage(ctx context.Context, id int) error
}

type languageUseCase struct {
	repository.LanguageRepository
}

func NewLanguageUseCase(r repository.LanguageRepository) LanguageUseCase {
	return &languageUseCase{r}
}

func (u *languageUseCase) GetLanguages(ctx context.Context) ([]*model.Language, error) {
	return u.LanguageRepository.GetLanguages(ctx)
}

func (u *languageUseCase) CreateLanguage(ctx context.Context, object *model.Language) (int, error) {
	return u.LanguageRepository.CreateLanguage(ctx, object)
}

func (u *languageUseCase) UpdateLanguage(ctx context.Context, object *model.Language) (bool, error) {
	return u.LanguageRepository.UpdateLanguage(ctx, object)
}

func (u *languageUseCase) DeleteLanguage(ctx context.Context, id int) error {
	return u.LanguageRepository.DeleteLanguage(ctx, id)
}
