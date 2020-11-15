package grpc

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/proto"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
)

// GetUserInfo get current user info using id or email (if email is not empty)
func (h *Server) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	var (
		usr entity.User
		err error
	)

	if in.Email != "" {
		usr, err = h.uc.GetUserByEmail(ctx, in.Email)
		if err != nil {
			return nil, err
		}
	} else {
		usr, err = h.uc.GetUserByID(ctx, in.UserID)
		if err != nil {
			return nil, err
		}
	}

	return &pb.GetUserInfoResp{
		UserID: usr.ID,
		Name:   usr.Name,
		Email:  usr.Email,
		JoinDate: &pb.Date{
			UnixTime: usr.CreatedTime.Unix(),
		},
		Status:        int32(usr.Status),
		XXX_sizecache: 0,
	}, nil
}

// CheckUserIsAuthenticated check the current authenticated user
func (h *Server) CheckUserIsAuthenticated(ctx context.Context, in *pb.CheckUserIsAuthenticatedReq) (*pb.CheckUserIsAuthenticatedRes, error) {
	var (
		resp = &pb.CheckUserIsAuthenticatedRes{
			IsAuthenticated: false,
		}
	)

	usr := struct {
		entity.User
		jwt.StandardClaims `json:"-"`
	}{}

	tkn, err := jwt.ParseWithClaims(in.AuthToken, &usr, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.secret), nil
	})

	if err == nil && tkn.Valid {
		resp.IsAuthenticated = true
		resp.User = &pb.AuthenticatedUser{
			UserID: usr.ID,
			Name:   usr.Name,
			Email:  usr.Email,
			JoinDate: &pb.Date{
				UnixTime: usr.CreatedTime.Unix(),
			},
			Status: int32(usr.Status),
		}
		return resp, nil
	}

	return resp, errors.New("authentication failed")
}

// CreateUser create new user in account database
func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	var (
		resp = &pb.CreateUserRes{}
	)

	err := s.uc.CreateNewUser(ctx, entity.ParamSaveUser{
		User: entity.User{
			Name:    in.GetName(),
			Email:   in.GetEmail(),
			PWDHash: in.GetPassword(),
		},
	})
	if err != nil {
		return resp, err
	}

	resp.Success = true
	return resp, nil
}

// GetAccessToken get access token from account
func (h *Server) GetAccessToken(ctx context.Context, in *pb.GetAccessTokenReq) (*pb.GetAccessTokenResp, error) {
	usr, err := h.uc.AuthenticateUser(ctx, in.Email, in.Password)
	if err != nil {
		return &pb.GetAccessTokenResp{}, err
	}

	expirationTime := time.Now().Add(6 * 60 * time.Minute)

	claims := struct {
		entity.User
		jwt.StandardClaims
	}{
		User: usr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString([]byte(h.secret))
	if err != nil {
		return &pb.GetAccessTokenResp{}, err
	}

	return &pb.GetAccessTokenResp{
		Token:     tokenString,
		ExpiredAt: claims.ExpiresAt,
	}, nil
}
