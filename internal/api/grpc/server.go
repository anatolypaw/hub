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
func (s *server) AddCodeForPrint(ctx context.Context, in *pb.AddCodeForPrintReq) (*pb.Empty, error) {
	err := s.mstore.AddCodeForPrint(ctx, in.Sname, in.Gtin, in.Serial, in.Crypto)
	return nil, err
}

// Возвращает КМ для нанесения
func (s *server) GetCodeForPrint(ctx context.Context, in *pb.GetCodeForPrintReq) (*pb.GetCodeForPrintResp, error) {
	code, err := s.mstore.GetCodeForPrint(ctx, in.Gtin, in.Tname, in.Proddate)

	return &pb.GetCodeForPrintResp{
		Gtin:   code.Gtin,
		Serial: code.Serial,
		Crypto: code.Crypto,
		Id:     code.PrintID,
	}, err
}

// Отмечает напечатанный код произведенным
func (s *server) ProducePrinted(ctx context.Context, in *pb.ProducePrintedReq) (*pb.Empty, error) {
	err := s.mstore.ProducePrinted(ctx, in.Tname, in.Gtin, in.Serial, in.Proddate)
	return &pb.Empty{}, err
}

// Отбраковывает код по его serial
func (s *server) DiscardBySerial(ctx context.Context, in *pb.DiscardBySerialReq) (*pb.Empty, error) {
	err := s.mstore.DiscardBySerial(ctx, in.Tname, in.Gtin, in.Serial)
	return &pb.Empty{}, err
}

// Возвращает количество произведенных на линии кодов
func (s *server) GetProducedCount(ctx context.Context, in *pb.GetProducedCountReq) (*pb.GetProducedCountResp, error) {
	count, err := s.mstore.GetProducedCount(ctx, in.Tname, in.Gtin, in.Date)
	return &pb.GetProducedCountResp{
		ThisTerm: count,
	}, err
}

// Добавляет продукт
func (s *server) AddGood(ctx context.Context, in *pb.AddGoodReq) (*pb.Empty, error) {
	err := s.mstore.AddGood(ctx, in.Sname, in.Gtin, in.Desc)
	return &pb.Empty{}, err
}

// Возвращает продукты и требуемое для них количество кодов
func (s *server) GetGoodsCodeReq(ctx context.Context, in *pb.Empty) (*pb.GetGoodsCodeResp, error) {
	goods, err := s.mstore.GetGoodsCodeReq(ctx)
	if err != nil {
		return &pb.GetGoodsCodeResp{}, err
	}

	var goods_resp []*pb.GetGoodsCodeGood
	for _, good := range goods {
		goods_resp = append(goods_resp,
			&pb.GetGoodsCodeGood{
				Gtin:  good.Gtin,
				Desc:  good.Desc,
				Count: good.Required,
			},
		)
	}

	return &pb.GetGoodsCodeResp{
		Good: goods_resp,
	}, nil
}

// Возвращает код на выгрузку
func (s *server) GetCodeForUpload(ctx context.Context, in *pb.GetCodeForUploadReq) (*pb.GetCodeForUploadResp, error) {
	code, err := s.mstore.GetCodeForUpload(ctx, in.Gtin)

	return &pb.GetCodeForUploadResp{
		Gtin:     code.Gtin,
		Serial:   code.Serial,
		Crypto:   code.Crypto,
		Proddate: code.Proddate,
		Entryid:  code.EntryID,
		Discard:  code.Discard,
	}, err
}

// Устанавливает код выгруженным
func (s *server) SetCodeUploaded(ctx context.Context, in *pb.SetCodeUploadedReq) (*pb.Empty, error) {
	return &pb.Empty{}, s.mstore.SetCodeUploaded(ctx, in.Gtin, in.Serial, in.Entryid)
}
