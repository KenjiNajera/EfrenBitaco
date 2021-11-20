package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type TyperequestUseCase interface {
	GetTyperequests(ctx context.Context) ([]*model.Typerequest, error)
	CreateTyperequest(ctx context.Context, object *model.Typerequest) (int, error)
	UpdateTyperequest(ctx context.Context, object *model.Typerequest) (bool, error)
	DeleteTyperequest(ctx context.Context, id int) error
}

type typerequestUseCase struct {
	repository.TyperequestRepository
}

func NewTyperequestUseCase(r repository.TyperequestRepository) TyperequestUseCase {
	return &typerequestUseCase{r}
}

func (u *typerequestUseCase) GetTyperequests(ctx context.Context) ([]*model.Typerequest, error) {
	return u.TyperequestRepository.GetTyperequests(ctx)
}

func (u *typerequestUseCase) CreateTyperequest(ctx context.Context, object *model.Typerequest) (int, error) {
	return u.TyperequestRepository.CreateTyperequest(ctx, object)
}

func (u *typerequestUseCase) UpdateTyperequest(ctx context.Context, object *model.Typerequest) (bool, error) {
	return u.TyperequestRepository.UpdateTyperequest(ctx, object)
}

func (u *typerequestUseCase) DeleteTyperequest(ctx context.Context, id int) error {
	return u.TyperequestRepository.DeleteTyperequest(ctx, id)
}
