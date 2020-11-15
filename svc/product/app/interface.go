package app

import (
	"context"

	"github.com/reyhanfahlevi/poc-test/svc/product/internal/entity"
)

type IProductUC interface {
	CreateNewProduct(ctx context.Context, product entity.ParamSaveProduct) error
	GetProductByShopID(ctx context.Context, shopID int64, limit, offset int) ([]entity.ProductDetail, error)
	GetProductByID(ctx context.Context, id int64) (entity.ProductDetail, error)
	UpdateProduct(ctx context.Context, data entity.ParamSaveProduct) error
	SaveProductImage(ctx context.Context, productID int64, images []entity.ProductImg) error
}
