package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/repo"
)

func TestNew(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name string
		args args
		want *UseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_AuthenticateUser(t *testing.T) {
	type fields struct {
		salt string
		r    repo.Repo
	}
	type args struct {
		ctx   context.Context
		email string
		pass  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				salt: tt.fields.salt,
				r:    tt.fields.r,
			}
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
	type fields struct {
		salt string
		r    repo.Repo
	}
	type args struct {
		ctx  context.Context
		user entity.ParamSaveUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				salt: tt.fields.salt,
				r:    tt.fields.r,
			}
			if err := u.CreateNewUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetUserByEmail(t *testing.T) {
	type fields struct {
		salt string
		r    repo.Repo
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				salt: tt.fields.salt,
				r:    tt.fields.r,
			}
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
	type fields struct {
		salt string
		r    repo.Repo
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				salt: tt.fields.salt,
				r:    tt.fields.r,
			}
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

func TestUseCase_getUser(t *testing.T) {
	type fields struct {
		salt string
		r    repo.Repo
	}
	type args struct {
		ctx     context.Context
		filters []entity.Filter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				salt: tt.fields.salt,
				r:    tt.fields.r,
			}
			got, err := u.getUser(tt.args.ctx, tt.args.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_passHash(t *testing.T) {
	type fields struct {
		salt string
		r    repo.Repo
	}
	type args struct {
		pass string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				salt: tt.fields.salt,
				r:    tt.fields.r,
			}
			if got := u.passHash(tt.args.pass); got != tt.want {
				t.Errorf("passHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
