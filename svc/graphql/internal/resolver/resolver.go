package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	accpb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/proto"
	productpb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/product/proto"
	shoppb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/shop/proto"
	"github.com/reyhanfahlevi/poc-test/svc/graphql/internal/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// AccountGRPCClient is the contract with gRPC client for Account service.
type AccountGRPCClient interface {
	GetUserInfo(ctx context.Context, in *accpb.GetUserInfoReq) (*accpb.GetUserInfoResp, error)
	CheckUserIsAuthenticated(ctx context.Context, in *accpb.CheckUserIsAuthenticatedReq) (*accpb.CheckUserIsAuthenticatedRes, error)
	GetAccessToken(ctx context.Context, in *accpb.GetAccessTokenReq) (*accpb.GetAccessTokenResp, error)
	CreateUser(ctx context.Context, in *accpb.CreateUserReq) (*accpb.CreateUserRes, error)
}

type ShopGRPCClient interface {
	GetShopInfo(ctx context.Context, in *shoppb.GetShopInfoReq) (*shoppb.GetShopInfoRes, error)
	RegisterShop(ctx context.Context, in *shoppb.RegisterShopReq) (*shoppb.RegisterShopResp, error)
}

type ProductGRPCClient interface {
	CreateProduct(ctx context.Context, in *productpb.CreateProductReq) (*productpb.CreateProductRes, error)
	UpdateProduct(ctx context.Context, in *productpb.UpdateProductReq) (*productpb.UpdateProductRes, error)
	UpdateImage(ctx context.Context, in *productpb.UpdateImageReq) (*productpb.UpdateImageRes, error)
	GetProductInfo(ctx context.Context, in *productpb.GetProductInfoReq) (*productpb.GetProductInfoRes, error)
	GetProductList(ctx context.Context, in *productpb.GetProductListReq) (*productpb.GetProductListRes, error)
}

// Resolver graphql resolver
type Resolver struct {
	account AccountGRPCClient
	shop    ShopGRPCClient
	product ProductGRPCClient
}

// New will instantiate the resolver
func New(account AccountGRPCClient, shop ShopGRPCClient, product ProductGRPCClient) *Resolver {
	return &Resolver{
		account: account,
		shop:    shop,
		product: product,
	}
}

func (r *Resolver) getAuthenticatedUser(ctx context.Context) (*model.User, error) {
	gc, err := r.ginContextFromContext(ctx)
	if err != nil {
		return &model.User{}, err
	}

	token := gc.GetHeader("AuthToken")
	resp, err := r.account.CheckUserIsAuthenticated(ctx, &accpb.CheckUserIsAuthenticatedReq{
		AuthToken: token,
	})
	if err != nil {
		return &model.User{}, errors.New("unauthorized")
	}

	return &model.User{
		ID:       resp.GetUser().GetUserID(),
		Name:     resp.GetUser().GetName(),
		Email:    resp.GetUser().GetEmail(),
		JoinDate: time.Unix(resp.GetUser().GetJoinDate().GetUnixTime(), 0).Format("2006-01-02 15:04:05"),
	}, nil
}

func (r *Resolver) ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return &gin.Context{}, err
	}
	return gc, nil
}

func (r *Resolver) getCurrentUserShopID(ctx context.Context) (int64, error) {
	user, err := r.getAuthenticatedUser(ctx)
	if err != nil {
		return 0, err
	}

	shop, err := r.shop.GetShopInfo(ctx, &shoppb.GetShopInfoReq{
		UserID: user.ID,
	})
	if err != nil {
		return 0, err
	}

	return shop.Id, nil
}
