package main

import (
	"flag"

	"github.com/fabrizioperria/toll/aggregator/aggregator"
	"github.com/fabrizioperria/toll/aggregator/transport"
)

func main() {
	httpListenAddr := flag.String("http", ":8081", "http listen address")
	grpcListenAddr := flag.String("grpc", ":8082", "grpc listen address")
	invoiceAggregator := aggregator.NewInvoiceAggregator()

	go transport.SetupGRPCTransport(*grpcListenAddr, invoiceAggregator)

	// time.Sleep(1 * time.Second)
	// cl := client.NewGRPCAggregatorClient("localhost:8082")
	// cl.Aggregate(types.Distance{
	// 	OBUID:     12345,
	// 	Value:     100,
	// 	Timestamp: time.Now().Unix(),
	// })

	transport.SetupHTTPTransport(*httpListenAddr, invoiceAggregator)
}
