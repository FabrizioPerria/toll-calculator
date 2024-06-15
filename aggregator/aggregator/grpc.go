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
		OBUID:     int(agg.ObuId),
		Value:     agg.Value,
		Timestamp: agg.Timestamp,
	})
}
