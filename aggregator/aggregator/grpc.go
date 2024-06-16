package aggregator

import (
	"context"
	"strconv"

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
		OBUID:     int(agg.ObuId),
		Value:     agg.Value,
		Timestamp: agg.Timestamp,
	})
}

func (s *GRPCAggregator) Invoice(ctx context.Context, inv *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	i, err := s.svc.GetInvoice(strconv.Itoa(int(inv.ObuId)))
	if err != nil {
		return nil, err
	}
	return &pb.InvoiceResponse{
		ObuId:     int32(i.ObuID),
		Amount:    i.Amount,
		Distance:  i.Distance,
		Timestamp: i.Timestamp,
	}, nil
}
