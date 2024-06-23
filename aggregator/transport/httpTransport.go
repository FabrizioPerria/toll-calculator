package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fabrizioperria/toll/aggregator/aggregator"
	"github.com/fabrizioperria/toll/shared/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupHTTPTransport(listenAddr string, agg aggregator.Aggregator) *http.Server {
	mux := http.NewServeMux()

	aggregateMetricHandler := NewHTTPMetricHandler("aggregate")
	aggregateHandler := handleAggregate(agg)
	aggregateHandler = aggregateMetricHandler.instrumentHandler(aggregateHandler)
	mux.HandleFunc("/aggregate", makeHTTPFunc(aggregateHandler))

	invoiceMetricHandler := NewHTTPMetricHandler("invoice")
	invoiceHandler := handleInvoice(agg)
	invoiceHandler = invoiceMetricHandler.instrumentHandler(invoiceHandler)
	mux.HandleFunc("/invoice", makeHTTPFunc(invoiceHandler))

	mux.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		if err == http.ErrServerClosed {
			// Server closed by Shutdown or Close
			return nil
		}
	}

	return nil
}

type APIError struct {
	Err        error
	StatusCode int
}

func (e APIError) Error() string {
	return e.Err.Error()
}

type HTTPFunc func(w http.ResponseWriter, r *http.Request) error

type HTTPMetricHandler struct {
	errCounter prometheus.Counter
}

func NewHTTPMetricHandler(reqName string) *HTTPMetricHandler {
	return &HTTPMetricHandler{
		errCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: fmt.Sprintf("http_%s_error_counter", reqName),
			Name:      "aggregator",
			Help:      "Total number of http " + reqName + " errors",
		}),
	}
}

func (h *HTTPMetricHandler) instrumentHandler(f HTTPFunc) HTTPFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := f(w, r)
		if err != nil {
			h.errCounter.Inc()
		}
		return err
	}
}

func makeHTTPFunc(f HTTPFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				http.Error(w, err.Error(), apiErr.StatusCode)
			}
		}
	}
}

func handleInvoice(agg aggregator.Aggregator) HTTPFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodGet {
			return APIError{StatusCode: http.StatusMethodNotAllowed, Err: fmt.Errorf("method not allowed")}
		}
		obuID := r.URL.Query().Get("obu_id")
		if obuID == "" {
			return APIError{StatusCode: http.StatusBadRequest, Err: fmt.Errorf("obu_id is required")}
		}

		res, err := agg.GetInvoice(obuID)
		if err != nil {
			return APIError{StatusCode: http.StatusInternalServerError, Err: err}
		}

		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(res)
	}
}

func handleAggregate(agg aggregator.Aggregator) HTTPFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodPost {
			return APIError{StatusCode: http.StatusMethodNotAllowed, Err: fmt.Errorf("method not allowed")}
		}
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			return APIError{StatusCode: http.StatusBadRequest, Err: fmt.Errorf("invalid request")}
		}
		if err := agg.Aggregate(distance); err != nil {
			return APIError{StatusCode: http.StatusInternalServerError, Err: err}
		}
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
