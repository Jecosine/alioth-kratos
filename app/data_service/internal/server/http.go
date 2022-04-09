package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gql "github.com/Jecosine/alioth-kratos/api/gql/v1"
	v1 "github.com/Jecosine/alioth-kratos/api/helloworld/v1"
	"github.com/Jecosine/alioth-kratos/app/data_service/graph"
	"github.com/Jecosine/alioth-kratos/app/data_service/graph/generated"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/conf"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/middleware"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	gHandler "github.com/gorilla/handlers"
)

func RegisterGraphQLServer(s *http.Server) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	s.Handle("/gql/query", srv)
	s.Handle("/gql", playground.Handler("Alioth GraphQL", "/gql/query"))
	//r.Get("/gql/ping", )
}
func WhitelistAuthRouters() selector.MatchFunc {
	whitelistRoutes := map[string]struct{}{
		"/api.auth.v1.Auth.Login":    {},
		"/api.auth.v1.Auth.Register": {},
	}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelistRoutes[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, cj *conf.Jwt, base *service.BaseService, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(middleware.NewJwtAuth(cj.Secret)).Match(WhitelistAuthRouters()).Build(),
		),
		http.Filter(gHandler.CORS(
			gHandler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			gHandler.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			gHandler.AllowedOrigins([]string{"*"}),
		)),
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
