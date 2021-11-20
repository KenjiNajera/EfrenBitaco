package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type ClientRepository interface {
	GetClients(ctx context.Context) ([]*model.Clients, error)
	GetClientByID(ctx context.Context, id int) (*model.ClientProyect, error)
	CreateClient(ctx context.Context, client *model.Client) (int, error)
	UpdateClient(ctx context.Context, client *model.Client) (bool, error)
	DeleteClient(ctx context.Context, id int) error
	Reactivate(ctx context.Context, id int) (bool, error)
}
