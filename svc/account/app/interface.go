package app

import (
	"context"

	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
)

// IAccountUC contract with UseCase
type IAccountUC interface {
	CreateNewUser(ctx context.Context, user entity.ParamSaveUser) error
	AuthenticateUser(ctx context.Context, email, pass string) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserByID(ctx context.Context, id int64) (entity.User, error)
}
