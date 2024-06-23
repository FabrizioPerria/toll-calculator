package client

import (
	"os"

	"github.com/fabrizioperria/toll/shared/types"
)

type AggregatorClient interface {
	Aggregate(distance types.Distance) error
	Invoice(obuID string) (types.Invoice, error)
}

func AggregatorClientFactory(aggClient string) AggregatorClient {
	switch aggClient {
	case "http":
		return NewHTTPAggregatorClient(os.Getenv("AGGREGATOR_HTTP_CLIENT"))
	case "grpc":
		return NewGRPCAggregatorClient(os.Getenv("AGGREGATOR_GRPC_CLIENT"))
	default:
		panic("Invalid aggregator client " + aggClient)
	}
}
