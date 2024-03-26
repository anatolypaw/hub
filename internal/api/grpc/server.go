package grpcapi

import (
	"context"
	pb "hub/internal/api/grpc/grpcapi"
	"hub/internal/mstore"
)

type server struct {
	pb.HubServer
	mstore *mstore.MStore
}

func New(mstore *mstore.MStore) server {
	return server{
		mstore: mstore,
	}
}

// Добавляет код для печати
func (s *server) AddCodeForPrint(ctx context.Context, in *pb.AddCodeForPrintReq) (*pb.EmptyResp, error) {
	err := s.mstore.AddCodeForPrint(ctx, in.Sname, in.Gtin, in.Serial, in.Crypto)
	return nil, err
}

// Возвращает КМ для нанесения
func (s *server) GetCodeForPrint(ctx context.Context, in *pb.GetCodeForPrintReq) (*pb.GetCodeForPrintResp, error) {
	code, err := s.mstore.GetCodeForPrint(ctx, in.Gtin, in.Tname)

	return &pb.GetCodeForPrintResp{
		Gtin:   code.Gtin,
		Serial: code.Serial,
		Crypto: code.Crypto,
		Id:     code.PrintID,
	}, err
}

// Отмечает напечатанный код произведенным
func (s *server) ProducePrinted(ctx context.Context, in *pb.ProducePrintedReq) (*pb.EmptyResp, error) {
	err := s.mstore.ProducePrinted(ctx, in.Tname, in.Gtin, in.Serial, in.Proddate)
	return &pb.EmptyResp{}, err
}

// Возвращает количество произведенных на линии кодов
func (s *server) GetProducedCount(ctx context.Context, in *pb.GetProducedCountReq) (*pb.Int64, error) {
	count, err := s.mstore.GetProducedCount(ctx, in.Tname, in.Gtin, in.Date)
	return &pb.Int64{
		Value: count,
	}, err
}

// Добавляет продукт
func (s *server) AddGood(ctx context.Context, in *pb.AddGoodReq) (*pb.EmptyResp, error) {
	err := s.mstore.AddGood(ctx, in.Sname, in.Gtin, in.Desc)
	return &pb.EmptyResp{}, err
}
