package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type AuthenticateRepository interface {
	Authenticate(ctx context.Context, login *model.Login) (*model.Authenticate, error)
	RecoveryPassword(ctx context.Context, email *model.RecoveryPassword) (bool, error)
	ChangePassword(ctx context.Context, obj *model.ChangePassword) (bool, error)
	VerifyAccount(ctx context.Context, obj *model.ChangePassword) (bool, error)
	AuthenticateMovil(ctx context.Context, obj *model.LoginMovil) (*model.AuthenticateMovil, error)
}
