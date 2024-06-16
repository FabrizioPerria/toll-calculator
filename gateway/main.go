package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/fabrizioperria/toll/aggregator/client"
	constants "github.com/fabrizioperria/toll/shared"
)

func main() {
	listenAddr := flag.String("listen-addr", ":8080", "server listen address")
	flag.Parse()
	// httpClient := client.NewHTTPAggregatorClient(constants.AggregatorHttpClient)
	grpcClient := client.NewGRPCAggregatorClient(constants.AggregatorGrpcClient)
	h := &invoiceHandler{
		client: grpcClient,
	}
	http.HandleFunc("/invoice", serveHTTP(h.handleGetInvoice))
	http.ListenAndServe(*listenAddr, nil)
}

type invoiceHandler struct {
	client client.AggregatorClient
}

func (h *invoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	obuId := r.URL.Query().Get("obu_id")
	log.Printf("Received request for invoice for OBU ID: %s", obuId)

	invoice, err := h.client.Invoice(obuId)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, invoice)
}

func writeJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func serveHTTP(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
