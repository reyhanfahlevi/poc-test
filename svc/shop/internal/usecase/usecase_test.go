package usecase

import (
	"context"
	_ "crypto/sha256"
	"reflect"
	"testing"

	"github.com/pkg/errors"

	"github.com/golang/mock/gomock"
	"github.com/reyhanfahlevi/poc-test/svc/shop/internal/entity"
)

func TestUseCase_CreateNewSeller(t *testing.T) {
	type args struct {
		ctx  context.Context
		shop entity.ParamSaveShop
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
				shop: entity.ParamSaveShop{
					Shop: entity.Shop{
						ID:      1,
						Name:    "John Thor",
						UserID:  1,
						Address: "Dimari",
						Status:  1,
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
				shop: entity.ParamSaveShop{
					Shop: entity.Shop{
						ID:      1,
						Name:    "John Thor",
						UserID:  1,
						Address: "Dimari",
						Status:  1,
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
			if err := u.CreateNewShop(tt.args.ctx, tt.args.shop); (err != nil) != tt.wantErr {
				t.Errorf("CreateNewShop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetSellerByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock *MockrepoInterface)
		want    entity.Shop
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, entity.ParamGetShop{
					Filters: []entity.Filter{{
						Field: "user_id",
						Value: a.userID,
					}},
					Limit:  1,
					Offset: 0,
					Sort:   entity.SortBy{},
				}).Return([]entity.Shop{
					{
						ID:      1,
						Name:    "John Thor",
						UserID:  1,
						Address: "Dimari",
						Status:  1,
					},
				}, nil)
			},
			want: entity.Shop{
				ID:      1,
				Name:    "John Thor",
				UserID:  1,
				Address: "Dimari",
				Status:  1,
			},
			wantErr: false,
		}, {
			name: "Test Error - Not Found",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, entity.ParamGetShop{
					Filters: []entity.Filter{{
						Field: "user_id",
						Value: a.userID,
					}},
					Limit:  1,
					Offset: 0,
					Sort:   entity.SortBy{},
				}).Return([]entity.Shop{}, nil)
			},
			want:    entity.Shop{},
			wantErr: true,
		}, {
			name: "Test Error - DB Error",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, entity.ParamGetShop{
					Filters: []entity.Filter{{
						Field: "user_id",
						Value: a.userID,
					}},
					Limit:  1,
					Offset: 0,
					Sort:   entity.SortBy{},
				}).Return([]entity.Shop{}, errors.New("some err"))
			},
			want:    entity.Shop{},
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
			got, err := u.GetShopByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSellerByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSellerByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetSellerByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Shop
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
				mock.EXPECT().Get(a.ctx, entity.ParamGetShop{
					Filters: []entity.Filter{{
						Field: "id",
						Value: a.id,
					}},
					Limit:  1,
					Offset: 0,
					Sort:   entity.SortBy{},
				}).Return([]entity.Shop{
					{
						ID:      1,
						Name:    "John Thor",
						UserID:  1,
						Address: "Dimari",
						Status:  1,
					},
				}, nil)
			},
			want: entity.Shop{
				ID:      1,
				Name:    "John Thor",
				UserID:  1,
				Address: "Dimari",
				Status:  1,
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
			got, err := u.GetShopByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShopByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShopByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
