package transport

import (
	"net"

	"github.com/fabrizioperria/toll/aggregator/aggregator"
	"github.com/fabrizioperria/toll/shared/types/pb"
	"google.golang.org/grpc"
)

func SetupGRPCTransport(listenAddr string, agg aggregator.Aggregator) {
	tcpListener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	defer tcpListener.Close()

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

	pb.RegisterAggregatorServiceServer(grpcServer, aggregator.NewGRPCAggregator(agg))

	grpcServer.Serve(tcpListener)
}
