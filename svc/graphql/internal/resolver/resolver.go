package resolver

import (
	"context"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/proto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// AccountGRPCClient is the gRPC client for Account service.
type AccountGRPCClient interface {
	GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error)
	CheckUserIsAuthenticated(ctx context.Context, in *pb.CheckUserIsAuthenticatedReq) (*pb.CheckUserIsAuthenticatedRes, error)
	GetAccessToken(ctx context.Context, in *pb.GetAccessTokenReq) (*pb.GetAccessTokenResp, error)
	CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error)
}

// Resolver graphql resolver
type Resolver struct {
	account AccountGRPCClient
}

// New will instantiate the resolver
func New(account AccountGRPCClient) *Resolver {
	return &Resolver{
		account: account,
	}
}
