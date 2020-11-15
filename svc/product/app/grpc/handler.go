package grpc

import (
	"context"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/product/proto"
	"github.com/reyhanfahlevi/poc-test/svc/product/internal/entity"
)

// GetProductInfo get current user info using id or email (if email is not empty)
func (h *Server) GetProductInfo(ctx context.Context, in *pb.GetProductInfoReq) (*pb.GetProductInfoRes, error) {
	var (
		product entity.ProductDetail
		err     error
	)

	product, err = h.uc.GetProductByID(ctx, in.GetProductID())
	if err != nil {
		return nil, err
	}

	imgs := []*pb.ProductImage{}
	for _, i := range product.Images {
		imgs = append(imgs, &pb.ProductImage{
			ProductID: i.ProductID,
			ImageURL:  i.ImageURL,
		})
	}

	return &pb.GetProductInfoRes{
		Product: &pb.ProductInfo{
			ProductID:   product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			ShopID:      product.ShopID,
			Images:      imgs,
		},
	}, nil
}

// CreateProduct create new user in product database
func (s *Server) CreateProduct(ctx context.Context, in *pb.CreateProductReq) (*pb.CreateProductRes, error) {
	var (
		resp = &pb.CreateProductRes{}
	)

	param := entity.ParamSaveProduct{
		Product: entity.Product{
			ShopID:      in.GetShopID(),
			Name:        in.GetName(),
			Description: in.GetDescription(),
			Price:       in.GetPrice(),
		},
		Images: []entity.ProductImg{},
	}

	for _, i := range in.GetImages() {
		param.Images = append(param.Images, entity.ProductImg{
			ImageURL: i.GetImageURL(),
		})
	}

	err := s.uc.CreateNewProduct(ctx, param)
	if err != nil {
		return resp, err
	}

	resp.Success = true
	return resp, nil
}

func (s *Server) UpdateProduct(ctx context.Context, in *pb.UpdateProductReq) (*pb.UpdateProductRes, error) {
	var (
		resp = &pb.UpdateProductRes{}
	)

	param := entity.ParamSaveProduct{
		Product: entity.Product{
			ID:          in.GetProductID(),
			ShopID:      in.GetShopID(),
			Name:        in.GetName(),
			Description: in.GetDescription(),
			Price:       in.GetPrice(),
			Status:      int(in.GetStatus()),
		},
		Images: []entity.ProductImg{},
	}

	for _, i := range in.GetImages() {
		param.Images = append(param.Images, entity.ProductImg{
			ImageURL: i.GetImageURL(),
		})
	}

	err := s.uc.UpdateProduct(ctx, param)
	if err != nil {
		return resp, err
	}

	resp.Success = true
	return resp, nil
}

func (s *Server) UpdateImage(ctx context.Context, in *pb.UpdateImageReq) (*pb.UpdateImageRes, error) {
	var (
		images = []entity.ProductImg{}
	)

	for _, i := range in.GetImages() {
		images = append(images, entity.ProductImg{
			ImageURL: i.GetImageURL(),
		})
	}

	err := s.uc.SaveProductImage(ctx, in.GetProductID(), images)
	if err != nil {
		return &pb.UpdateImageRes{}, err
	}

	return &pb.UpdateImageRes{
		Success: true,
	}, nil
}

func (s *Server) GetProductList(ctx context.Context, in *pb.GetProductListReq) (*pb.GetProductListRes, error) {
	var (
		resp = &pb.GetProductListRes{
			Products: []*pb.ProductInfo{},
		}
	)
	products, err := s.uc.GetProductByShopID(ctx, in.GetShopID(), int(in.GetLimit()), int(in.GetOffset()))
	if err != nil {
		return &pb.GetProductListRes{}, nil
	}

	for _, p := range products {
		imgs := []*pb.ProductImage{}
		for _, i := range p.Images {
			imgs = append(imgs, &pb.ProductImage{
				ProductID: i.ProductID,
				ImageURL:  i.ImageURL,
			})
		}

		resp.Products = append(resp.Products, &pb.ProductInfo{
			ProductID:   p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			ShopID:      p.ShopID,
			Images:      imgs,
		})
	}

	return resp, nil
}
