package server

import (
	"github.com/Jecosine/alioth-kratos/internal/server/judger"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)

// JudgerProviderSet is server providers for judger.
var JudgerProviderSet = wire.NewSet(judger.NewJudgerHTTPServer, judger.NewJudgerGRPCServer)
