package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type TyperequestRepository interface {
	GetTyperequests(ctx context.Context) ([]*model.Typerequest, error)
	CreateTyperequest(ctx context.Context, object *model.Typerequest) (int, error)
	UpdateTyperequest(ctx context.Context, object *model.Typerequest) (bool, error)
	DeleteTyperequest(ctx context.Context, id int) error
}
