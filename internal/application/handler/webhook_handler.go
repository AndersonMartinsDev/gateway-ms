package handler

import (
	"gateway-ms/internal/application/service"
	"gateway-ms/internal/infrastructure/response"
	"io"
	"net/http"
)

// UserHandler é o handler HTTP para as requisições de usuário.
type WebhookHandler struct {
	webhookService *service.WebhookService
}

// NewUserHandler cria uma nova instância de UserHandler.
func NewWebhookHandler(s *service.WebhookService) *WebhookHandler {
	return &WebhookHandler{
		webhookService: s,
	}
}

func (handler WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo", http.StatusInternalServerError)
		return
	}

	// Chama o método do serviço para processar e publicar o webhook
	if err := handler.webhookService.ProcessWebhook(r.Context(), body); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	res := struct {
		Data string `json:"data"`
	}{
		Data: "Mensagem na fila de processanto ",
	}
	response.Response(w, http.StatusOK, res, nil)
}
