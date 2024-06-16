package main

import (
	"flag"

	"github.com/fabrizioperria/toll/aggregator/aggregator"
	"github.com/fabrizioperria/toll/aggregator/transport"
)

func main() {
	httpListenAddr := flag.String("http", ":8081", "http listen address")
	grpcListenAddr := flag.String("grpc", ":8082", "grpc listen address")
	flag.Parse()
	invoiceAggregator := aggregator.NewInvoiceAggregator()

	go transport.SetupGRPCTransport(*grpcListenAddr, invoiceAggregator)
	transport.SetupHTTPTransport(*httpListenAddr, invoiceAggregator)
}
