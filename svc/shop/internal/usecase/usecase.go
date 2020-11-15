package usecase

import (
	"context"

	"github.com/reyhanfahlevi/poc-test/svc/shop/internal/entity"
)

type repoInterface interface {
	Get(ctx context.Context, param entity.ParamGetShop) ([]entity.Shop, error)
	Save(ctx context.Context, data entity.ParamSaveShop) error
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

// CreateNewShop create new shop
func (u *UseCase) CreateNewShop(ctx context.Context, shop entity.ParamSaveShop) error {

	return u.r.Save(ctx, shop)
}

func (u *UseCase) GetShopByUserID(ctx context.Context, userID int64) (entity.Shop, error) {
	return u.getShop(ctx, []entity.Filter{
		{
			Field: "user_id",
			Value: userID,
		},
	})
}

func (u *UseCase) GetShopByID(ctx context.Context, id int64) (entity.Shop, error) {
	return u.getShop(ctx, []entity.Filter{
		{
			Field: "id",
			Value: id,
		},
	})
}

func (u *UseCase) getShop(ctx context.Context, filters []entity.Filter) (entity.Shop, error) {
	shops, err := u.r.Get(ctx, entity.ParamGetShop{
		Filters: filters,
		Limit:   1,
		Offset:  0,
		Sort:    entity.SortBy{},
	})
	if err != nil {
		return entity.Shop{}, err
	}

	if len(shops) < 1 {
		return entity.Shop{}, entity.ErrNotFound
	}

	return shops[0], nil
}
