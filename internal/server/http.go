package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gql "github.com/Jecosine/alioth-kratos/api/gql/v1"
	v1 "github.com/Jecosine/alioth-kratos/api/helloworld/v1"
	"github.com/Jecosine/alioth-kratos/graph"
	"github.com/Jecosine/alioth-kratos/graph/generated"
	"github.com/Jecosine/alioth-kratos/internal/conf"
	"github.com/Jecosine/alioth-kratos/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func RegisterGraphQLServer(s *http.Server) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	s.Handle("/gql/query", srv)
	s.Handle("/gql", playground.Handler("Alioth GraphQL", "/gql/query"))
	//r.Get("/gql/ping", )
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, base *service.BaseService, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	gql.RegisterBaseHTTPServer(srv, base)
	RegisterGraphQLServer(srv)
	return srv
}
