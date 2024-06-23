package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/fabrizioperria/toll/aggregator/client"
	"github.com/fabrizioperria/toll/shared/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()

	l := logger.LoggerFactory(os.Getenv("LOG_PATH"))

	h := &invoiceHandler{
		client: aggregatorClientFactory(os.Getenv("AGGREGATOR_CLIENT")),
	}
	http.HandleFunc("/invoice", serveHTTP(h.handleGetInvoice, l))
	http.ListenAndServe(os.Getenv("GATEWAY_ENDPOINT"), nil)
}

func aggregatorClientFactory(aggClient string) client.AggregatorClient {
	switch aggClient {
	case "grpc":
		return client.NewGRPCAggregatorClient(os.Getenv("AGGREGATOR_GRPC_ENDPOINT"))
	case "http":
		return client.NewHTTPAggregatorClient(os.Getenv("AGGREGATOR_HTTP_ENDPOINT"))
	default:
		panic("invalid aggregator client")
	}
}

type invoiceHandler struct {
	client client.AggregatorClient
}

func (h *invoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	obuId := r.URL.Query().Get("obu_id")
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

func serveHTTP(fn apiFunc, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			logger.WithFields(logrus.Fields{
				"method":   r.Method,
				"uri":      r.RequestURI,
				"duration": float64(time.Since(start).Milliseconds()) / 1000,
			}).Info("Request processed")
		}(time.Now())

		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"uri":    r.RequestURI,
				"error":  err.Error(),
			}).Error("Request failed")
		}
	}
}
