package handler

import (
	"gateway-ms/internal/application/service"
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
	// 1. Ler o corpo (payload) da requisição HTTP
	// var payload interface{}

	// err := request.Serialization(r.Body, &payload)
	// if err != nil {
	// 	response.Erro(w, http.StatusBadRequest, err)
	// 	return
	// }
	// defer r.Body.Close()

	// payloadBytes, _ := json.Marshal(payload)
	// // 2. Criar a mensagem gRPC com o payload lido e um ID único
	// req := &proto.ProcessWebhookRequest{
	// 	Id:      uuid.NewString(),     // Gera um ID único para rastreamento
	// 	Payload: string(payloadBytes), // Converte os bytes do payload para string
	// }

	// defer r.Body.Close()

	// 3. Chamar o serviço (o restante da sua lógica)
	// ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) // Aumentado o timeout para dar tempo de processar
	// defer cancel()

	// res, err := handler.webhookService.Process(ctx, req)
	// if err != nil {
	// 	slog.Error("Erro ao chamar o serviço de webhook", "error", err)
	// 	http.Error(w, "Erro ao processar o webhook.", http.StatusInternalServerError)
	// 	return
	// }

	// slog.Info("Resposta do Webhook Processor", "status", res.GetStatus(), "message", res.GetMessage())
	// w.Write([]byte("Webhook processado com sucesso!"))
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo", http.StatusInternalServerError)
		return
	}

	// Chama o método do serviço para processar e publicar o webhook
	if err := handler.webhookService.ProcessWebhook(r.Context(), body); err != nil {
		http.Error(w, "Erro ao processar o webhook", http.StatusInternalServerError)
		return
	}

	// Responde OK imediatamente, sem esperar pelo processamento.
	w.WriteHeader(http.StatusOK)
}
