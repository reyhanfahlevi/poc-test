package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/reyhanfahlevi/poc-test/pkg/config"
	accClient "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/account/client"
	productClient "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/product/client"
	shopClient "github.com/reyhanfahlevi/poc-test/pkg/grpc-client/shop/client"
	"github.com/reyhanfahlevi/poc-test/svc/graphql/internal/generated"
	"github.com/reyhanfahlevi/poc-test/svc/graphql/internal/resolver"
)

// Defining the Graphql handler
func graphqlHandler(h *handler.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	err := config.Init(config.WithConfigFile("files/etc/graphql/config.json", "/etc/graphql/config.json"))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Get()

	ac, err := accClient.GetClient(accClient.Options{Address: cfg.GRPCServer.Account})
	if err != nil {
		log.Fatal(err)
	}

	sc, err := shopClient.GetClient(shopClient.Options{Address: cfg.GRPCServer.Shop})
	if err != nil {
		log.Fatal(err)
	}

	pc, err := productClient.GetClient(productClient.Options{Address: cfg.GRPCServer.Product})
	if err != nil {
		log.Fatal(err)
	}

	rslv := resolver.New(ac, sc, pc)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: rslv}))

	r := gin.Default()
	r.Use(ginContextToContextMiddleware())
	r.POST("/query", graphqlHandler(srv))
	r.GET("/", playgroundHandler())

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}

	port := cfg.Server.Port
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
