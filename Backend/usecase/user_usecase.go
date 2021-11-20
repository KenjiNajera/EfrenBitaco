package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *model.User) (bool, error)
	UpdateUser(ctx context.Context, user *model.User) (bool, error)
	ActivateUser(ctx context.Context, id int) (bool, error)
	DesactivateUser(ctx context.Context, id int) (bool, error)
	GetUsers(ctx context.Context) ([]*model.GetUserDatas, error)
	GetMyData(ctx context.Context, id int) (*model.MyData, error)
	GetMyDataMovil(ctx context.Context, id int) (*model.MyData, error)
	UpdateMyData(ctx context.Context, user *model.MyData) (bool, error)
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (bool, error) {
	return u.UserRepository.CreateUser(ctx, user)
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) (bool, error) {
	return u.UserRepository.UpdateUser(ctx, user)
}

func (u *userUseCase) ActivateUser(ctx context.Context, id int) (bool, error) {
	return u.UserRepository.ActivateUser(ctx, id)
}

func (u *userUseCase) DesactivateUser(ctx context.Context, id int) (bool, error) {
	return u.UserRepository.DesactivateUser(ctx, id)
}

func (u *userUseCase) GetUsers(ctx context.Context) ([]*model.GetUserDatas, error) {
	return u.UserRepository.GetUsers(ctx)
}

func (u *userUseCase) GetMyData(ctx context.Context, id int) (*model.MyData, error) {
	return u.UserRepository.GetMyData(ctx, id)
}

func (u *userUseCase) GetMyDataMovil(ctx context.Context, id int) (*model.MyData, error) {
	return u.UserRepository.GetMyDataMovil(ctx, id)
}

func (u *userUseCase) UpdateMyData(ctx context.Context, user *model.MyData) (bool, error) {
	return u.UserRepository.UpdateMyData(ctx, user)
}
