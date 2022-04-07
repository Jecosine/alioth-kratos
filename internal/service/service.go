package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewBaseService)
var JudgerProviderSet = wire.NewSet(NewJudgerService)
