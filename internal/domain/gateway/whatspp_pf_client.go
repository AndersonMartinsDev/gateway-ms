package gateway

import (
	"context"
	"gateway-ms/proto"
)

type WhatsappPFClient interface {
	GetPairQrCodeConnect(ctx context.Context, req *proto.GenerateQRCodeRequest) (*proto.GenerateQRCodeResponse, error)
}
