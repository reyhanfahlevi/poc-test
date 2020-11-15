package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	accpb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/proto"
	productpb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/product/proto"
	shoppb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/shop/proto"
	"github.com/reyhanfahlevi/poc-test/svc/graphql/internal/generated"
	"github.com/reyhanfahlevi/poc-test/svc/graphql/internal/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (bool, error) {
	resp, err := r.account.CreateUser(ctx, &accpb.CreateUserReq{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	})
	if err != nil {
		return false, err
	}

	return resp.GetSuccess(), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	resp, err := r.account.GetAccessToken(ctx, &accpb.GetAccessTokenReq{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return "", err
	}

	return resp.GetToken(), nil
}

func (r *mutationResolver) OpenShop(ctx context.Context, input model.OpenShop) (bool, error) {
	user, err := r.getAuthenticatedUser(ctx)
	if err != nil {
		return false, err
	}

	resp, err := r.shop.RegisterShop(ctx, &shoppb.RegisterShopReq{
		UserID:  user.ID,
		Name:    input.Name,
		Address: input.Address,
	})
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

func (r *mutationResolver) AddProduct(ctx context.Context, input *model.NewProduct) (bool, error) {
	shopID, err := r.getCurrentUserShopID(ctx)
	if err != nil {
		return false, err
	}
	imgs := []*productpb.ProductImage{}
	for _, i := range input.Images {
		imgs = append(imgs, &productpb.ProductImage{ImageURL: i})
	}
	resp, err := r.product.CreateProduct(ctx, &productpb.CreateProductReq{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		ShopID:      shopID,
		Images:      imgs,
	})
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

func (r *mutationResolver) EditProduct(ctx context.Context, input *model.EditProduct) (bool, error) {
	shopID, err := r.getCurrentUserShopID(ctx)
	if err != nil {
		return false, err
	}

	product, err := r.product.GetProductInfo(ctx, &productpb.GetProductInfoReq{ProductID: input.ProductID})
	if err != nil {
		return false, err
	}

	if product.GetProduct().GetShopID() != shopID {
		return false, errors.New("that product is belong someone else")
	}

	imgs := []*productpb.ProductImage{}
	for _, i := range input.Images {
		imgs = append(imgs, &productpb.ProductImage{ImageURL: i})
	}

	resp, err := r.product.UpdateProduct(ctx, &productpb.UpdateProductReq{
		ProductID:   input.ProductID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		ShopID:      shopID,
		Status:      1,
		Images:      imgs,
	})
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, productID int64) (bool, error) {
	shopID, err := r.getCurrentUserShopID(ctx)
	if err != nil {
		return false, err
	}

	product, err := r.product.GetProductInfo(ctx, &productpb.GetProductInfoReq{ProductID: productID})
	if err != nil {
		return false, err
	}

	if product.GetProduct().GetShopID() != shopID {
		return false, errors.New("that product is belong someone else")
	}

	resp, err := r.product.UpdateProduct(ctx, &productpb.UpdateProductReq{
		ProductID:   product.GetProduct().GetProductID(),
		Name:        product.GetProduct().GetName(),
		Description: product.GetProduct().GetDescription(),
		Price:       product.GetProduct().GetPrice(),
		ShopID:      product.GetProduct().GetShopID(),
		Images:      product.GetProduct().GetImages(),
		Status:      0,
	})
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	resp, err := r.account.GetUserInfo(ctx, &accpb.GetUserInfoReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	joinDate := time.Unix(resp.GetJoinDate().GetUnixTime(), 0)

	return &model.User{
		ID:       resp.GetUserID(),
		Name:     resp.GetName(),
		Email:    resp.GetEmail(),
		JoinDate: joinDate.Format("2006-01-02 15:04:05"),
	}, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	resp, err := r.account.GetUserInfo(ctx, &accpb.GetUserInfoReq{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}

	joinDate := time.Unix(resp.GetJoinDate().GetUnixTime(), 0)

	return &model.User{
		ID:       resp.GetUserID(),
		Name:     resp.GetName(),
		Email:    resp.GetEmail(),
		JoinDate: joinDate.Format("2006-01-02 15:04:05"),
	}, nil
}

func (r *queryResolver) GetShopInfo(ctx context.Context, shopID int64) (*model.Shop, error) {
	shop, err := r.shop.GetShopInfo(ctx, &shoppb.GetShopInfoReq{
		ShopID: shopID,
	})
	if err != nil {
		return &model.Shop{}, err
	}

	return &model.Shop{
		ID:     shop.GetId(),
		Name:   shop.GetName(),
		Status: int64(shop.GetStatus()),
	}, nil
}

func (r *queryResolver) GetMyShopInfo(ctx context.Context) (*model.Shop, error) {
	user, err := r.getAuthenticatedUser(ctx)
	if err != nil {
		return &model.Shop{}, err
	}

	shop, err := r.shop.GetShopInfo(ctx, &shoppb.GetShopInfoReq{
		UserID: user.ID,
	})
	if err != nil {
		return &model.Shop{}, err
	}

	return &model.Shop{
		ID:      shop.GetId(),
		Name:    shop.GetName(),
		Address: shop.GetAddress(),
		Status:  int64(shop.GetStatus()),
	}, nil
}

func (r *queryResolver) GetProductInfo(ctx context.Context, productID int64) (*model.Product, error) {
	resp, err := r.product.GetProductInfo(ctx, &productpb.GetProductInfoReq{ProductID: productID})
	if err != nil {
		return &model.Product{}, err
	}

	imgs := []*model.ProductImage{}
	for _, i := range resp.GetProduct().GetImages() {
		imgs = append(imgs, &model.ProductImage{
			ProductID: i.GetProductID(),
			ImageURL:  i.GetImageURL(),
		})
	}

	return &model.Product{
		ProductID:   resp.GetProduct().GetProductID(),
		ShopID:      resp.GetProduct().GetShopID(),
		Name:        resp.GetProduct().GetName(),
		Description: resp.GetProduct().GetDescription(),
		Price:       resp.GetProduct().GetPrice(),
		Images:      imgs,
		Status:      int64(resp.GetProduct().GetStatus()),
	}, nil
}

func (r *queryResolver) GetProductList(ctx context.Context, limit int64, offset int64) ([]*model.Product, error) {
	var (
		products = []*model.Product{}
	)

	shopID, err := r.getCurrentUserShopID(ctx)
	if err != nil {
		return products, err
	}

	resp, err := r.product.GetProductList(ctx, &productpb.GetProductListReq{
		Limit:  int32(limit),
		Offset: int32(offset),
		ShopID: shopID,
	})
	if err != nil {
		return products, err
	}

	for _, p := range resp.GetProducts() {
		imgs := []*model.ProductImage{}

		for _, i := range p.GetImages() {
			imgs = append(imgs, &model.ProductImage{
				ProductID: i.GetProductID(),
				ImageURL:  i.GetImageURL(),
			})
		}

		products = append(products, &model.Product{
			ProductID:   p.GetProductID(),
			ShopID:      p.GetShopID(),
			Name:        p.GetName(),
			Description: p.GetDescription(),
			Price:       p.GetPrice(),
			Images:      imgs,
			Status:      int64(p.GetStatus()),
		})
	}

	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
