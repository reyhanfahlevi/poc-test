package grpc

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/reyhanfahlevi/pkg/go/log"
	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/product/proto"
	"github.com/reyhanfahlevi/poc-test/svc/product/app"
	"github.com/tokopedia/panics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server struct
type Server struct {
	server *grpc.Server
	uc     app.IProductUC
	secret string
}

// New instantiate grpc server
func New(uc app.IProductUC) *Server {
	grpcSvr := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			panics.UnaryServerInterceptor,
		)),
	)

	srv := &Server{
		server: grpcSvr,
		uc:     uc,
	}

	pb.RegisterProductServer(srv.server, srv)
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

//GracefulStop gracefully stop server
func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}
