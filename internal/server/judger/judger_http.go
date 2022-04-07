package judger

import (
	v1 "github.com/Jecosine/alioth-kratos/api/judger/v1"
	"github.com/Jecosine/alioth-kratos/internal/conf"
	"github.com/Jecosine/alioth-kratos/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewJudgerHTTPServer new an HTTP server.
func NewJudgerHTTPServer(c *conf.Server, judger *service.JudgerService, logger log.Logger) *http.Server {
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
	v1.RegisterJudgerHTTPServer(srv, judger)
	return srv
}
