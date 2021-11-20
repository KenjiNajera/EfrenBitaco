package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type RequestUseCase interface {
	GetRequests(ctx context.Context) ([]*model.GetRequest, error)
	CreateRequest(ctx context.Context, request *model.Request) (bool, error)
	ResponseRequest(ctx context.Context, request *model.Request) (bool, error)
	GetRequestsUser(ctx context.Context, id int) ([]*model.GetRequestUser, error)
}

type requestUseCase struct {
	repository.RequestRepository
}

func NewRequestUseCase(r repository.RequestRepository) RequestUseCase {
	return &requestUseCase{r}
}

func (u *requestUseCase) GetRequests(ctx context.Context) ([]*model.GetRequest, error) {
	return u.RequestRepository.GetRequests(ctx)
}

func (u *requestUseCase) CreateRequest(ctx context.Context, request *model.Request) (bool, error) {
	return u.RequestRepository.CreateRequest(ctx, request)
}

func (u *requestUseCase) ResponseRequest(ctx context.Context, request *model.Request) (bool, error) {
	return u.RequestRepository.ResponseRequest(ctx, request)
}

func (u *requestUseCase) GetRequestsUser(ctx context.Context, id int) ([]*model.GetRequestUser, error) {
	return u.RequestRepository.GetRequestsUser(ctx, id)
}
