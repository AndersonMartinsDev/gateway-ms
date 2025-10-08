package grpc_client

import (
	"context"
	"gateway-ms/internal/domain/gateway"
	"gateway-ms/proto"

	"google.golang.org/grpc"
)

type WhatsappPFClient struct {
	client proto.WhatsappServiceClient
}

func NewWhatsappPFClient(conn *grpc.ClientConn) gateway.WhatsappPFClient {
	return &WhatsappPFClient{
		client: proto.NewWhatsappServiceClient(conn),
	}
}

func (c *WhatsappPFClient) GetPairQrCodeConnect(ctx context.Context, req *proto.GenerateQRCodeRequest) (*proto.GenerateQRCodeResponse, error) {
	return c.client.GenerateQRCode(ctx, req)
}
