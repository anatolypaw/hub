package grpcapi

import (
	"context"
	pb "hub/internal/api/grpc/grpcapi"
	"hub/internal/service/produce"
)

type server struct {
	pb.HubServer
	produce *produce.Produce
}

func New(produce *produce.Produce) server {
	return server{produce: produce}
}

// Возвращает КМ для нанесения
func (s *server) GetCodeForPrint(ctx context.Context, in *pb.GetCodeForPrintRequest) (*pb.Code, error) {
	code, err := s.produce.GetCodeForPrint(ctx, in.GetGtin(), in.GetTname())

	return &pb.Code{
		Gtin:   code.Gtin,
		Serial: code.Serial,
		Crypto: code.Crypto,
	}, err
}
