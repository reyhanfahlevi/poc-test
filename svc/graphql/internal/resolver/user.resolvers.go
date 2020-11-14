package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/proto"
	"github.com/reyhanfahlevi/poc-test/svc/graphql/internal/generated"
	"github.com/reyhanfahlevi/poc-test/svc/graphql/internal/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (bool, error) {
	resp, err := r.account.CreateUser(ctx, &pb.CreateUserReq{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	})
	if err != nil {
		return false, err
	}

	return resp.GetSuccess(), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	resp, err := r.account.GetAccessToken(ctx, &pb.GetAccessTokenReq{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return "", nil
	}

	return resp.GetToken(), nil
}

func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	resp, err := r.account.GetUserInfo(ctx, &pb.GetUserInfoReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	joinDate := time.Unix(resp.GetJoinDate().GetUnixTime(), 0)

	return &model.User{
		ID:       resp.GetUserID(),
		Name:     resp.GetName(),
		Email:    resp.GetEmail(),
		JoinDate: joinDate.Format("2006-01-02 15:04:05"),
	}, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	resp, err := r.account.GetUserInfo(ctx, &pb.GetUserInfoReq{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}

	joinDate := time.Unix(resp.GetJoinDate().GetUnixTime(), 0)

	return &model.User{
		ID:       resp.GetUserID(),
		Name:     resp.GetName(),
		Email:    resp.GetEmail(),
		JoinDate: joinDate.Format("2006-01-02 15:04:05"),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
