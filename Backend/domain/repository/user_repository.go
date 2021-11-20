package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (bool, error)
	UpdateUser(ctx context.Context, user *model.User) (bool, error)
	ActivateUser(ctx context.Context, id int) (bool, error)
	DesactivateUser(ctx context.Context, id int) (bool, error)
	GetUsers(ctx context.Context) ([]*model.GetUserDatas, error)
	GetMyData(ctx context.Context, id int) (*model.MyData, error) // Pendiente
	GetMyDataMovil(ctx context.Context, id int) (*model.MyData, error)
	UpdateMyData(ctx context.Context, user *model.MyData) (bool, error) // Pendiente
}
