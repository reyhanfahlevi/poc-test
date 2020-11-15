package usecase

import (
	"context"

	"github.com/reyhanfahlevi/pkg/go/log"

	"github.com/reyhanfahlevi/poc-test/svc/product/internal/entity"
)

type repoInterface interface {
	Get(ctx context.Context, param entity.ParamGetProduct) ([]entity.Product, error)
	GetProductImage(ctx context.Context, productID int64) ([]entity.ProductImg, error)
	Save(ctx context.Context, data entity.ParamSaveProduct) error
	UpdateProduct(ctx context.Context, param entity.ParamSaveProduct) error
	UpdateProductImage(ctx context.Context, productID int64, images []entity.ProductImg) error
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

// CreateNewProduct create new product
func (u *UseCase) CreateNewProduct(ctx context.Context, product entity.ParamSaveProduct) error {
	return u.r.Save(ctx, product)
}

func (u *UseCase) GetProductByShopID(ctx context.Context, shopID int64, limit, offset int) ([]entity.ProductDetail, error) {
	var (
		productsDetail = []entity.ProductDetail{}
		param          = entity.ParamGetProduct{
			Filters: []entity.Filter{
				{
					Field: "shop_id",
					Value: shopID,
				},
			},
			Limit:  limit,
			Offset: offset,
		}
	)

	if param.Limit > 25 {
		param.Limit = 25
	}

	products, err := u.r.Get(ctx, param)
	if err != nil {
		return productsDetail, err
	}

	for _, p := range products {
		tmp := entity.ProductDetail{
			Product: p,
			Images:  u.getProductImg(ctx, p.ID),
		}

		productsDetail = append(productsDetail, tmp)
	}

	return productsDetail, nil
}

// for now ignore any error when fetching images
func (u *UseCase) getProductImg(ctx context.Context, productID int64) []entity.ProductImg {
	images, err := u.r.GetProductImage(ctx, productID)
	if err != nil {
		log.Error(err)
	}

	return images
}

func (u *UseCase) GetProductByID(ctx context.Context, id int64) (entity.ProductDetail, error) {
	products, err := u.r.Get(ctx, entity.ParamGetProduct{
		Filters: []entity.Filter{
			{
				Field: "id",
				Value: id,
			},
		},
		Limit:  1,
		Offset: 0,
		Sort:   entity.SortBy{},
	})
	if err != nil {
		return entity.ProductDetail{}, err
	}

	if len(products) < 1 {
		return entity.ProductDetail{}, entity.ErrNotFound
	}

	return entity.ProductDetail{
		Product: products[0],
		Images:  u.getProductImg(ctx, products[0].ID),
	}, nil
}

func (u *UseCase) SaveProductImage(ctx context.Context, productID int64, images []entity.ProductImg) error {
	return u.r.UpdateProductImage(ctx, productID, images)
}

func (u *UseCase) UpdateProduct(ctx context.Context, data entity.ParamSaveProduct) error {
	return u.r.UpdateProduct(ctx, data)
}
