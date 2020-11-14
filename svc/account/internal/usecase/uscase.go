package usecase

import (
	"context"
	"crypto"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"

	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
)

type repoInterface interface {
	Get(ctx context.Context, param entity.ParamGetUser) ([]entity.User, error)
	Save(ctx context.Context, data entity.ParamSaveUser) error
}

// UseCase struct
type UseCase struct {
	salt string
	r    repoInterface
}

// New instantiate usecase instance
func New(salt string, r repoInterface) *UseCase {
	return &UseCase{
		salt: salt,
		r:    r,
	}
}

// CreateNewUser create new user
func (u *UseCase) CreateNewUser(ctx context.Context, user entity.ParamSaveUser) error {
	user.PWDHash = u.passHash(user.PWDHash)
	return u.r.Save(ctx, user)
}

// AuthenticateUser auth user by email and pass
func (u *UseCase) AuthenticateUser(ctx context.Context, email, pass string) (entity.User, error) {
	return u.getUser(ctx, []entity.Filter{
		{
			Field: "email",
			Value: email,
		}, {
			Field: "pwd_hash",
			Value: u.passHash(pass),
		},
	})
}

func (u *UseCase) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return u.getUser(ctx, []entity.Filter{
		{
			Field: "email",
			Value: email,
		},
	})
}

func (u *UseCase) GetUserByID(ctx context.Context, id int64) (entity.User, error) {
	return u.getUser(ctx, []entity.Filter{
		{
			Field: "id",
			Value: id,
		},
	})
}

func (u *UseCase) getUser(ctx context.Context, filters []entity.Filter) (entity.User, error) {
	users, err := u.r.Get(ctx, entity.ParamGetUser{
		Filters: filters,
		Limit:   1,
		Offset:  0,
		Sort:    entity.SortBy{},
	})
	if err != nil {
		return entity.User{}, err
	}

	if len(users) < 1 {
		return entity.User{}, errors.New("user not found")
	}

	return users[0], nil
}

func (u *UseCase) passHash(pass string) string {
	sha := crypto.SHA256.New()
	sha.Reset()

	// ignore the error because the func never return error
	_, _ = sha.Write([]byte(fmt.Sprintf("%s%s", u.salt, pass)))
	return hex.EncodeToString(sha.Sum(nil))
}
