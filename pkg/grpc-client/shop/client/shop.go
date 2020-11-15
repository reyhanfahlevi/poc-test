package client

import (
	"context"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/shop/proto"
)

// GetShopInfo will get shop info
func (c *Client) GetShopInfo(ctx context.Context, in *pb.GetShopInfoReq) (*pb.GetShopInfoRes, error) {
	return c.client.GetShopInfo(ctx, in)
}

// RegisterShop create a shop for user
func (c *Client) RegisterShop(ctx context.Context, in *pb.RegisterShopReq) (*pb.RegisterShopResp, error) {
	return c.client.RegisterShop(ctx, in)
}
