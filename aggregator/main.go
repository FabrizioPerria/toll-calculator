package main

import (
	"flag"

	"github.com/fabrizioperria/toll/aggregator/aggregator"
	"github.com/fabrizioperria/toll/aggregator/transport"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8081", "http listen address")
	invoiceAggregator := aggregator.NewInvoiceAggregator()

	transport.SetupHTTPTransport(*listenAddr, invoiceAggregator)
}
