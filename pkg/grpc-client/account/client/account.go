package client

import (
	"context"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/proto"
)

// GetUserInfo get the current user info from account service
func (c *Client) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	return c.client.GetUserInfo(ctx, in)
}

// CheckIsUserAuthenticated check current authenticated user by authToken
func (c *Client) CheckIsUserAuthenticated(ctx context.Context, in *pb.CheckUserIsAuthenticatedReq) (*pb.CheckUserIsAuthenticatedRes, error) {
	return c.client.CheckUserIsAuthenticated(ctx, in)
}

// GetAccessToken get user access token by email and password
func (c *Client) GetAccessToken(ctx context.Context, in *pb.GetAccessTokenReq) (*pb.GetAccessTokenResp, error) {
	return c.client.GetAccessToken(ctx, in)
}

// CreateUser create new user
func (c *Client) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	return c.client.CreateUser(ctx, in)
}
