package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/repository"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type AuthenticateUseCase interface {
	Authenticate(ctx context.Context, login *model.Login) (*model.Authenticate, error)
	RecoveryPassword(ctx context.Context, email *model.RecoveryPassword) (bool, error)
	ChangePassword(ctx context.Context, obj *model.ChangePassword) (bool, error)
	VerifyAccount(ctx context.Context, obj *model.ChangePassword) (bool, error)
	AuthenticateMovil(ctx context.Context, obj *model.LoginMovil) (*model.AuthenticateMovil, error)
}

type authenticateUseCase struct {
	repository.AuthenticateRepository
}

func NewAuthenticateUseCase(r repository.AuthenticateRepository) AuthenticateUseCase {
	return &authenticateUseCase{r}
}

func (u *authenticateUseCase) Authenticate(ctx context.Context, login *model.Login) (*model.Authenticate, error) {
	return u.AuthenticateRepository.Authenticate(ctx, login)
}

func (u *authenticateUseCase) RecoveryPassword(ctx context.Context, email *model.RecoveryPassword) (bool, error) {
	return u.AuthenticateRepository.RecoveryPassword(ctx, email)
}

func (u *authenticateUseCase) ChangePassword(ctx context.Context, obj *model.ChangePassword) (bool, error) {
	return u.AuthenticateRepository.ChangePassword(ctx, obj)
}

func (u *authenticateUseCase) AuthenticateMovil(ctx context.Context, obj *model.LoginMovil) (*model.AuthenticateMovil, error) {
	return u.AuthenticateRepository.AuthenticateMovil(ctx, obj)
}

func (u *authenticateUseCase) VerifyAccount(ctx context.Context, obj *model.ChangePassword) (bool, error) {
	return u.AuthenticateRepository.VerifyAccount(ctx, obj)
}
