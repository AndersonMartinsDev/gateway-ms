package handler

import (
	"fmt"
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/infrastructure/response"
	"net/http"
)

// UserHandler é o handler HTTP para as requisições de usuário.
type WhatsappPFHandler struct {
	service *service.WhatsAppPFService
}

// NewUserHandler cria uma nova instância de UserHandler.
func NewWhatsappPFHandler(s *service.WhatsAppPFService) *WhatsappPFHandler {
	return &WhatsappPFHandler{
		service: s,
	}
}

func (handler WhatsappPFHandler) HandleQrCode(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phoneNumber")
	qrcode, err := handler.service.GetPairQrCodeConnect(phoneNumber)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, fmt.Errorf("error: %w", err))
		return
	}

	response.JSON(w, http.StatusAccepted, qrcode)
}
