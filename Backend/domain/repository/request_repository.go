package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type RequestRepository interface {
	GetRequests(ctx context.Context) ([]*model.GetRequest, error)
	CreateRequest(ctx context.Context, request *model.Request) (bool, error)
	ResponseRequest(ctx context.Context, request *model.Request) (bool, error)
	GetRequestsUser(ctx context.Context, id int) ([]*model.GetRequestUser, error)
}
