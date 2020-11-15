package usecase

import (
	"context"
	_ "crypto/sha256"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/reyhanfahlevi/poc-test/svc/product/internal/entity"
)

func TestUseCase_CreateNewProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		product entity.ParamSaveProduct
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock *MockrepoInterface)
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx: context.Background(),
				product: entity.ParamSaveProduct{
					Product: entity.Product{
						ID:          1,
						Name:        "Baju Bajuri",
						ShopID:      1,
						Description: "Mulus like new",
						Status:      1,
					},
					Images: []entity.ProductImg{
						{ImageURL: "https://via.placeholder.com/300"},
						{ImageURL: "https://via.placeholder.com/400"},
					},
				},
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Save(a.ctx, gomock.Any()).Return(nil)
			},
			wantErr: false,
		}, {
			name: "Test Err",
			args: args{
				ctx: context.Background(),
				product: entity.ParamSaveProduct{
					Product: entity.Product{
						ID:          1,
						Name:        "Baju Bajuri",
						ShopID:      1,
						Description: "Mulus like new",
						Status:      1,
					},
					Images: []entity.ProductImg{
						{ImageURL: "https://via.placeholder.com/300"},
						{ImageURL: "https://via.placeholder.com/400"},
					},
				},
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Save(a.ctx, gomock.Any()).Return(entity.ErrIsExist)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockrepoInterface(ctrl)

			u := New("secret", mockRepo)
			tt.mock(tt.args, mockRepo)
			if err := u.CreateNewProduct(tt.args.ctx, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("CreateNewProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetProductByShopID(t *testing.T) {
	type args struct {
		ctx    context.Context
		shopID int64
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock *MockrepoInterface)
		want    []entity.ProductDetail
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx:    context.Background(),
				shopID: 1,
				limit:  26,
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, entity.ParamGetProduct{
					Filters: []entity.Filter{{
						Field: "shop_id",
						Value: a.shopID,
					}},
					Limit:  a.limit - 1,
					Offset: a.offset,
				}).Return([]entity.Product{
					{
						ID:          1,
						Name:        "John Thor",
						ShopID:      1,
						Description: "Dimari",
						Status:      1,
					},
				}, nil)

				mock.EXPECT().GetProductImage(a.ctx, gomock.Any()).Return([]entity.ProductImg{}, nil)
			},
			want: []entity.ProductDetail{
				{
					Product: entity.Product{
						ID:          1,
						Name:        "John Thor",
						ShopID:      1,
						Description: "Dimari",
						Status:      1,
					},
					Images: []entity.ProductImg{},
				},
			},
			wantErr: false,
		}, {
			name: "Test Error - DB Error",
			args: args{
				ctx:    context.Background(),
				shopID: 1,
				limit:  1,
				offset: 0,
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, entity.ParamGetProduct{
					Filters: []entity.Filter{{
						Field: "shop_id",
						Value: a.shopID,
					}},
					Limit:  a.limit,
					Offset: a.offset,
					Sort:   entity.SortBy{},
				}).Return([]entity.Product{}, errors.New("some err"))
			},
			want:    []entity.ProductDetail{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockrepoInterface(ctrl)

			u := New("secret", mockRepo)
			tt.mock(tt.args, mockRepo)
			got, err := u.GetProductByShopID(tt.args.ctx, tt.args.shopID, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetProductByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		want    entity.ProductDetail
		mock    func(a args, mock *MockrepoInterface)
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, entity.ParamGetProduct{
					Filters: []entity.Filter{{
						Field: "id",
						Value: a.id,
					}},
					Limit:  1,
					Offset: 0,
					Sort:   entity.SortBy{},
				}).Return([]entity.Product{
					{
						ID:          1,
						Name:        "John Thor",
						ShopID:      1,
						Description: "Dimari",
						Status:      1,
					},
				}, nil)

				mock.EXPECT().GetProductImage(a.ctx, gomock.Any()).Return([]entity.ProductImg{}, nil)
			},
			want: entity.ProductDetail{
				Product: entity.Product{
					ID:          1,
					Name:        "John Thor",
					ShopID:      1,
					Description: "Dimari",
					Status:      1,
				},
				Images: []entity.ProductImg{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := NewMockrepoInterface(ctrl)

			u := New("secret", mockRepo)
			tt.mock(tt.args, mockRepo)
			got, err := u.GetProductByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
