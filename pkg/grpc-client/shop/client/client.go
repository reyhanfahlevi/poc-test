package client

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	pb "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/shop/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// Client struct
type Client struct {
	client pb.ShopClient
	opts   Options
	conn   *grpc.ClientConn

	// redial lock protect us from doing method call when we do redial
	redialLock sync.RWMutex
}

// Options struct define option(s) available for affiliate grpc client
type Options struct {
	Address string
}

const (
	defaultConnectTimeout = 10 * time.Second
)

// GetClient create & returns reference of client struct
func GetClient(o Options) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultConnectTimeout)
	defer cancel()
	return GetClientWithContext(ctx, o)
}

// GetClientWithContext creates & returns reference of client struct that implement Client interface with given context
func GetClientWithContext(ctx context.Context, o Options) (*Client, error) {
	c := &Client{
		opts: o,
	}
	return c, c.dial(ctx)
}

func (c *Client) dial(ctx context.Context) error {
	conn, err := grpc.DialContext(ctx, c.opts.Address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(c.redialInterceptor),
	)
	if err != nil {
		log.Println("[ERROR-GRPC] error dialing to:", c.opts.Address, err)
		if conn != nil {
			conn.Close()
		}
		return err
	}
	c.conn = conn

	// wait for it to be connected
	// we wait without returns error because while failed wait means the server is down,
	// we still hope for it to be up again in the near future.
	numTries := 0
	for state := conn.GetState(); state != connectivity.Ready; state = conn.GetState() {
		curCtx, _ := context.WithTimeout(ctx, 4*time.Second)
		if conn.WaitForStateChange(curCtx, state) != true {
			log.Println("[ERROR-GRPC] failed in waitforstatechange")
		}

		numTries++
		if numTries > 5 {
			log.Println("[ERROR-GRPC] failed in waitforstatechange - max attempts exceeded to", c.opts.Address, state)
			break
		}
	}

	c.client = pb.NewShopClient(conn)
	return nil
}

// redialInterceptor is interceptor which check if we need to  redial the connection.
// we need to redial in this condition:
// - transport error
func (c *Client) redialInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// invokes the method.
	// we make it as anonymous func protected by read lock for these reasons:
	// - we don't make a call when we do redial
	// - unlock guaranteed to be called in case of `panic`
	// - we will reuse it again, below
	invoke := func() error {
		c.redialLock.RLock()
		defer c.redialLock.RUnlock()
		return invoker(ctx, method, req, reply, cc, opts...)
	}

	// invoke the method and return in case of no error or not transport error
	err := invoke()
	if err == nil || !c.isTransportError(err) {
		return err
	}

	// redial
	err = func() error {
		// protect it with write lock, so there is no grpc call
		// when we do redial
		c.redialLock.Lock()
		defer c.redialLock.Unlock()

		log.Println("redialing...")

		oldConn := c.conn

		ctx, cancel := context.WithTimeout(context.Background(), defaultConnectTimeout)
		defer cancel()
		err := c.dial(ctx)
		if err == nil && oldConn != nil {
			oldConn.Close() // remove old connection, because we already have the new one
		}
		return err
	}()

	if err != nil {
		return err
	}

	// re-invoke
	return invoke()
}

// isTransportError check if the error we get is transport error
func (c *Client) isTransportError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "transport is closing") ||
		strings.Contains(errStr, "client connection is closing") ||
		strings.Contains(errStr, "no connection available")
}
