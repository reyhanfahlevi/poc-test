package client

import (
	"context"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/product/proto"
)

func (c *Client) CreateProduct(ctx context.Context, in *pb.CreateProductReq) (*pb.CreateProductRes, error) {
	return c.client.CreateProduct(ctx, in)
}

func (c *Client) UpdateProduct(ctx context.Context, in *pb.UpdateProductReq) (*pb.UpdateProductRes, error) {
	return c.client.UpdateProduct(ctx, in)
}

func (c *Client) UpdateImage(ctx context.Context, in *pb.UpdateImageReq) (*pb.UpdateImageRes, error) {
	return c.client.UpdateImage(ctx, in)
}

func (c *Client) GetProductInfo(ctx context.Context, in *pb.GetProductInfoReq) (*pb.GetProductInfoRes, error) {
	return c.client.GetProductInfo(ctx, in)
}

func (c *Client) GetProductList(ctx context.Context, in *pb.GetProductListReq) (*pb.GetProductListRes, error) {
	return c.client.GetProductList(ctx, in)
}
