package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fabrizioperria/toll/shared/types"
)

type HTTPAggregatorClient struct {
	endpoint string
}

func NewHTTPAggregatorClient(endpoint string) AggregatorClient {
	return &HTTPAggregatorClient{endpoint: endpoint}
}

func (c *HTTPAggregatorClient) Aggregate(distance types.Distance) error {
	distanceMarshal, err := json.Marshal(distance)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.endpoint+"/aggregate", bytes.NewReader(distanceMarshal))
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	return nil
}

func (c *HTTPAggregatorClient) Invoice(obuID string) (types.Invoice, error) {
	req, err := http.NewRequest("GET", c.endpoint+"/invoice?obu_id="+obuID, nil)
	if err != nil {
		return types.Invoice{}, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return types.Invoice{}, err
	}

	if response.StatusCode != http.StatusOK {
		return types.Invoice{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var invoice types.Invoice
	err = json.NewDecoder(response.Body).Decode(&invoice)
	if err != nil {
		return types.Invoice{}, err
	}
	return invoice, nil
}
