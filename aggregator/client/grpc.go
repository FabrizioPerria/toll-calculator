package client

import (
	"context"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/fabrizioperria/toll/shared/types/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCAggregatorClient struct {
	client   pb.AggregatorServiceClient
	endpoint string
}

func NewGRPCAggregatorClient(endpoint string) AggregatorClient {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &GRPCAggregatorClient{
		endpoint: endpoint,
		client:   pb.NewAggregatorServiceClient(conn),
	}
}

func (c *GRPCAggregatorClient) Aggregate(distance types.Distance) error {
	_, err := c.client.Aggregate(context.Background(), &pb.AggregateRequest{
		ObuId:     distance.ObuId,
		Value:     distance.Value,
		Timestamp: distance.Timestamp,
	})
	return err
}

func (c *GRPCAggregatorClient) Invoice(obuID string) (types.Invoice, error) {
	inv, err := c.client.Invoice(context.Background(), &pb.InvoiceRequest{
		ObuId: obuID,
	})
	if err != nil {
		return types.Invoice{}, err
	}
	return types.Invoice{
		ObuId:     inv.ObuId,
		Amount:    inv.Amount,
		Distance:  inv.Distance,
		Timestamp: inv.Timestamp,
	}, nil
}
