package repo

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
)

func getMockDB() (*sqlx.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock
	}

	return sqlx.NewDb(mockDB, "postgres"), mock
}

func TestRepo_Get(t *testing.T) {
	mockTime := time.Now()
	type args struct {
		ctx   context.Context
		param entity.ParamGetUser
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock sqlmock.Sqlmock)
		want    []entity.User
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx: context.Background(),
				param: entity.ParamGetUser{
					Filters: []entity.Filter{
						{
							Field: "id",
							Value: 1,
						},
					},
					Limit:  1,
					Offset: 0,
				},
			},
			mock: func(arg args, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT (.+) FROM wp_user WHERE id = (.+) LIMIT (.+) OFFSET (.+)").
					WithArgs(arg.param.Filters[0].Value, arg.param.Limit, arg.param.Offset).
					WillReturnRows(
						sqlmock.NewRows(
							[]string{"id", "name", "email", "pwd_hash", "status", "created_time", "created_by", "updated_time", "updated_by"}).
							AddRow(1, "Test", "john.thor@wp.com", "12312123", 1, mockTime, 0, mockTime, 0).
							AddRow(2, "Test 2", "odading@wp.com", "12312123", 1, mockTime, 0, mockTime, 0),
					)
			},
			want: []entity.User{
				{
					ID:          1,
					Name:        "Test",
					Email:       "john.thor@wp.com",
					PWDHash:     "12312123",
					Status:      1,
					CreatedTime: mockTime,
					UpdatedTime: mockTime,
				}, {
					ID:          2,
					Name:        "Test 2",
					Email:       "odading@wp.com",
					PWDHash:     "12312123",
					Status:      1,
					CreatedTime: mockTime,
					UpdatedTime: mockTime,
				},
			},
			wantErr: false,
		}, {
			name: "Test Failed",
			args: args{
				ctx: context.Background(),
				param: entity.ParamGetUser{
					Filters: []entity.Filter{
						{
							Field: "id",
							Value: 1,
						},
					},
					Limit:  1,
					Offset: 0,
				},
			},
			mock: func(arg args, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT (.+) FROM wp_user WHERE id = (.+) LIMIT (.+) OFFSET (.+)").
					WithArgs(arg.param.Filters[0].Value, arg.param.Limit, arg.param.Offset).
					WillReturnError(errors.New("expect error"))
			},
			want:    []entity.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := getMockDB()
			r := New(db)

			tt.mock(tt.args, mock)
			got, err := r.Get(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_Save(t *testing.T) {
	type args struct {
		ctx  context.Context
		data entity.ParamSaveUser
	}
	tests := []struct {
		name    string
		args    args
		mock    func(a args, mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "Test Success",
			args: args{
				ctx: context.Background(),
				data: entity.ParamSaveUser{
					User: entity.User{
						Name:    "Test",
						Email:   "john.thor@wp.com",
						PWDHash: "hashhh",
					},
				},
			},
			mock: func(a args, mock sqlmock.Sqlmock) {
				q := `INSERT INTO wp_user (.+) VALUES (.+) RETURNING id`
				mock.ExpectQuery(q).
					WithArgs(a.data.Name, a.data.Email, a.data.PWDHash, a.data.CreatedBy, a.data.CreatedBy).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: false,
		}, {
			name: "Test Error - Conflict",
			args: args{
				ctx: context.Background(),
				data: entity.ParamSaveUser{
					User: entity.User{
						Name:    "Test",
						Email:   "john.thor@wp.com",
						PWDHash: "hashhh",
					},
				},
			},
			mock: func(a args, mock sqlmock.Sqlmock) {
				q := `INSERT INTO wp_user (.+) VALUES (.+) RETURNING id`
				mock.ExpectQuery(q).
					WithArgs(a.data.Name, a.data.Email, a.data.PWDHash, a.data.CreatedBy, a.data.CreatedBy).
					WillReturnError(errors.New("pq: duplicate key value violates unique constraint"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := getMockDB()
			r := New(db)

			tt.mock(tt.args, mock)
			if err := r.Save(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
