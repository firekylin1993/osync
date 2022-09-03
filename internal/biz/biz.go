package biz

import (
	"github.com/google/wire"
	"osync/internal/biz/bizagg"
	"osync/internal/biz/bizedn"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	bizagg.NewAggUsecase,
	bizedn.NewEdnUsecase,
)
