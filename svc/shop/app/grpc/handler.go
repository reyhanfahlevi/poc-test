package grpc

import (
	"context"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/shop/proto"
	"github.com/reyhanfahlevi/poc-test/svc/shop/internal/entity"
)

// GetShopInfo get current user info using id or email (if email is not empty)
func (h *Server) GetShopInfo(ctx context.Context, in *pb.GetShopInfoReq) (*pb.GetShopInfoRes, error) {
	var (
		shop entity.Shop
		err  error
	)

	if in.UserID != 0 {
		shop, err = h.uc.GetShopByUserID(ctx, in.UserID)
		if err != nil {
			return nil, err
		}
	} else {
		shop, err = h.uc.GetShopByID(ctx, in.ShopID)
		if err != nil {
			return nil, err
		}
	}

	return &pb.GetShopInfoRes{
		Id:            shop.ID,
		UserID:        shop.ID,
		Name:          shop.Name,
		Address:       shop.Address,
		Status:        int32(shop.Status),
		XXX_sizecache: 0,
	}, nil
}

func (s *Server) RegisterShop(ctx context.Context, in *pb.RegisterShopReq) (*pb.RegisterShopResp, error) {
	var (
		resp = &pb.RegisterShopResp{}
	)

	err := s.uc.CreateNewShop(ctx, entity.ParamSaveShop{
		Shop: entity.Shop{
			Name:    in.GetName(),
			UserID:  in.GetUserID(),
			Address: in.GetAddress(),
		},
	})
	if err != nil {
		return resp, err
	}

	resp.Success = true
	return resp, nil
}
