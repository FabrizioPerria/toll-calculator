package main

import (
	"os"

	"github.com/fabrizioperria/toll/aggregator/aggregator"
	"github.com/fabrizioperria/toll/aggregator/transport"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	httpListenAddr := os.Getenv("AGGREGATOR_HTTP_ENDPOINT")
	grpcListenAddr := os.Getenv("AGGREGATOR_GRPC_ENDPOINT")

	invoiceAggregator := aggregator.NewInvoiceAggregator()

	go transport.SetupGRPCTransport(grpcListenAddr, invoiceAggregator)
	transport.SetupHTTPTransport(httpListenAddr, invoiceAggregator)
}
