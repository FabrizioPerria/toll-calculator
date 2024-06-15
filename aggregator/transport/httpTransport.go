package transport

import (
	"encoding/json"
	"net/http"

	"github.com/fabrizioperria/toll/aggregator/aggregator"
	"github.com/fabrizioperria/toll/shared/types"
)

func SetupHTTPTransport(listenAddr string, agg aggregator.Aggregator) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/aggregate", handleAggregate(agg))
	mux.HandleFunc("/invoice", handleInvoice(agg))

	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		if err == http.ErrServerClosed {
			// Server closed by Shutdown or Close
			return nil
		}
	}

	return nil
}

func handleInvoice(agg aggregator.Aggregator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		obuID := r.URL.Query().Get("obu_id")
		if obuID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing obu_id"))
			return
		}

		res, err := agg.GetInvoice(obuID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func handleAggregate(agg aggregator.Aggregator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := agg.Aggregate(distance); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
