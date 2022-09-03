package data

import (
	"github.com/google/wire"
	"osync/internal/data/myagg"
	"osync/internal/data/myedn"
	"osync/internal/data/myotel"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	myotel.NewTracerClient,
	myotel.NewTracerExporter,
	myotel.NewMetricClient,
	myotel.NewMetricExporter,

	myagg.NewAggRepo,
	myedn.NewEdnRepo,
)
