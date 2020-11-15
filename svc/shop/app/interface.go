package app

import (
	"context"

	"github.com/reyhanfahlevi/poc-test/svc/shop/internal/entity"
)

// IAccountUC contract with UseCase
type ISellerUC interface {
	CreateNewShop(ctx context.Context, user entity.ParamSaveShop) error
	GetShopByUserID(ctx context.Context, id int64) (entity.Shop, error)
	GetShopByID(ctx context.Context, id int64) (entity.Shop, error)
}
