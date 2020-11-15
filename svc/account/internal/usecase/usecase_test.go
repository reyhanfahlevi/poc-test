package usecase

import (
	"context"
	_ "crypto/sha256"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
)

func TestUseCase_AuthenticateUser(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
		pass  string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock *MockrepoInterface)
		want    entity.User
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx:   context.Background(),
				email: "john.thor@wp.com",
				pass:  "hayo",
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, gomock.Any()).Return([]entity.User{
					{
						ID:      1,
						Name:    "John Thor",
						Email:   "john.thor@wp.com",
						PWDHash: "asrasrar",
						Status:  1,
					},
				}, nil)
			},
			want: entity.User{
				ID:      1,
				Name:    "John Thor",
				Email:   "john.thor@wp.com",
				PWDHash: "asrasrar",
				Status:  1,
			},
			wantErr: false,
		}, {
			name: "Test Err - Not Found",
			args: args{
				ctx:   context.Background(),
				email: "john.thor@wp.com",
				pass:  "hayo",
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, gomock.Any()).Return([]entity.User{}, nil)
			},
			want:    entity.User{},
			wantErr: true,
		}, {
			name: "Test Err - DB Error",
			args: args{
				ctx:   context.Background(),
				email: "john.thor@wp.com",
				pass:  "hayo",
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, gomock.Any()).Return([]entity.User{}, errors.New("some err"))
			},
			want:    entity.User{},
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
			got, err := u.AuthenticateUser(tt.args.ctx, tt.args.email, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_CreateNewUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user entity.ParamSaveUser
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
				user: entity.ParamSaveUser{
					User: entity.User{
						ID:      1,
						Name:    "test",
						Email:   "test@wp.com",
						PWDHash: "asdas",
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
				user: entity.ParamSaveUser{
					User: entity.User{
						ID:      1,
						Name:    "test",
						Email:   "test@wp.com",
						PWDHash: "asdas",
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
			if err := u.CreateNewUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock *MockrepoInterface)
		want    entity.User
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx:   context.Background(),
				email: "john.thor@wp.com",
			},
			mock: func(a args, mock *MockrepoInterface) {
				mock.EXPECT().Get(a.ctx, entity.ParamGetUser{
					Filters: []entity.Filter{{
						Field: "email",
						Value: a.email,
					}},
					Limit:  1,
					Offset: 0,
					Sort:   entity.SortBy{},
				}).Return([]entity.User{
					{
						ID:      1,
						Name:    "John Thor",
						Email:   "john.thor@wp.com",
						PWDHash: "asrasrar",
						Status:  1,
					},
				}, nil)
			},
			want: entity.User{
				ID:      1,
				Name:    "John Thor",
				Email:   "john.thor@wp.com",
				PWDHash: "asrasrar",
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
			got, err := u.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetUserByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		want    entity.User
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
				mock.EXPECT().Get(a.ctx, entity.ParamGetUser{
					Filters: []entity.Filter{{
						Field: "id",
						Value: a.id,
					}},
					Limit:  1,
					Offset: 0,
					Sort:   entity.SortBy{},
				}).Return([]entity.User{
					{
						ID:      1,
						Name:    "John Thor",
						Email:   "john.thor@wp.com",
						PWDHash: "asrasrar",
						Status:  1,
					},
				}, nil)
			},
			want: entity.User{
				ID:      1,
				Name:    "John Thor",
				Email:   "john.thor@wp.com",
				PWDHash: "asrasrar",
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
			got, err := u.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
