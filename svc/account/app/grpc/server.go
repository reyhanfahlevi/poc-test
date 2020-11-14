package grpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/reyhanfahlevi/pkg/go/log"

	"github.com/dgrijalva/jwt-go"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/proto"
	"github.com/reyhanfahlevi/poc-test/svc/account/app"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
	"github.com/tokopedia/panics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server struct
type Server struct {
	server *grpc.Server
	uc     app.IAccountUC
	secret string
}

// New instantiate grpc server
func New(secret string, uc app.IAccountUC) *Server {
	grpcSvr := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			panics.UnaryServerInterceptor,
		)),
	)

	srv := &Server{
		server: grpcSvr,
		uc:     uc,
		secret: secret,
	}

	pb.RegisterAccountServer(srv.server, srv)
	reflection.Register(srv.server)

	return srv
}

func (s *Server) Run(port string) error {
	idleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		s.GracefulStop()
		log.Info("GRPC server shutdown gracefully")
		close(idleConnsClosed)
	}()

	log.Info("GRPC server running on port", port)
	if err := s.Serve(port); err != http.ErrServerClosed {
		// Error starting or closing listener:
		return err
	}

	<-idleConnsClosed
	log.Info("GRPC server stopping")
	return nil
}

// Serve grpc
func (s *Server) Serve(port string) error {
	lis, err := net.Listen("tcp4", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}

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
	}

	return resp, err
}

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

//GracefulStop gracefully stop server
func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}
