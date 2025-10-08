package service

import (
	"context"
	"gateway-ms/internal/domain/gateway"
	"gateway-ms/proto"
)

type WhatsAppPFService struct {
	client gateway.WhatsappPFClient
}

func NewWhatsAppPFService(client gateway.WhatsappPFClient) *WhatsAppPFService {
	return &WhatsAppPFService{
		client: client,
	}
}

func (w *WhatsAppPFService) GetPairQrCodeConnect(phoneNumber string) (string, error) {
	context := context.Background()

	res, err := w.client.GetPairQrCodeConnect(context, &proto.GenerateQRCodeRequest{
		PhoneNumber: phoneNumber,
	})
	return res.QrCode, err
}
