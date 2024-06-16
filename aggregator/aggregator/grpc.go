package aggregator

import (
	"context"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/fabrizioperria/toll/shared/types/pb"
)

type GRPCAggregator struct {
	pb.UnimplementedAggregatorServiceServer
	svc Aggregator
}

func NewGRPCAggregator(svc Aggregator) *GRPCAggregator {
	return &GRPCAggregator{svc: svc}
}

func (s *GRPCAggregator) Aggregate(ctx context.Context, agg *pb.AggregateRequest) (*pb.None, error) {
	return nil, s.svc.Aggregate(types.Distance{
		ObuId:     agg.ObuId,
		Value:     agg.Value,
		Timestamp: agg.Timestamp,
	})
}

func (s *GRPCAggregator) Invoice(ctx context.Context, inv *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	i, err := s.svc.GetInvoice(inv.ObuId)
	if err != nil {
		return nil, err
	}
	return &pb.InvoiceResponse{
		ObuId:     i.ObuId,
		Amount:    i.Amount,
		Distance:  i.Distance,
		Timestamp: i.Timestamp,
	}, nil
}
