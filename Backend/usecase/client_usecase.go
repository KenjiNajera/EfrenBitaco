package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
)

type ClientUseCase interface {
	GetClients(ctx context.Context) ([]*model.Clients, error)
	GetClientByID(ctx context.Context, id int) (*model.ClientProyect, error)
	CreateClient(ctx context.Context, client *model.Client) (int, error)
	UpdateClient(ctx context.Context, client *model.Client) (bool, error)
	DeleteClient(ctx context.Context, id int) error
	Reactivate(ctx context.Context, id int) (bool, error)
}

type clientUseCase struct {
	repository.ClientRepository
}

func NewClientUseCase(r repository.ClientRepository) ClientUseCase {
	return &clientUseCase{r}
}

func (u *clientUseCase) GetClients(ctx context.Context) ([]*model.Clients, error) {
	return u.ClientRepository.GetClients(ctx)
}
func (u *clientUseCase) GetClientByID(ctx context.Context, id int) (*model.ClientProyect, error) {
	return u.ClientRepository.GetClientByID(ctx, id)
}
func (u *clientUseCase) CreateClient(ctx context.Context, client *model.Client) (int, error) {
	return u.ClientRepository.CreateClient(ctx, client)
}
func (u *clientUseCase) UpdateClient(ctx context.Context, client *model.Client) (bool, error) {
	return u.ClientRepository.UpdateClient(ctx, client)
}
func (u *clientUseCase) DeleteClient(ctx context.Context, id int) error {
	return u.ClientRepository.DeleteClient(ctx, id)
}
func (u *clientUseCase) Reactivate(ctx context.Context, id int) (bool, error) {
	return u.ClientRepository.Reactivate(ctx, id)
}
