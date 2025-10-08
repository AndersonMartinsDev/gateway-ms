package service

import (
	"context"
	"gateway-ms/internal/domain/message"
	"log/slog"
)

// WebhookService orquestra o processamento do webhook.
type WebhookService struct {
	publisher message.MessagePublisher
}

// NewWebhookService cria uma nova instância de WebhookService.
func NewWebhookService(publisher message.MessagePublisher) *WebhookService {
	return &WebhookService{
		publisher: publisher,
	}
}

// ProcessWebhook recebe o payload bruto do webhook
// e o publica em uma fila para processamento assíncrono.
func (s *WebhookService) ProcessWebhook(ctx context.Context, payload []byte) error {
	slog.Info("Recebido novo webhook, publicando na fila...")

	// O nome da fila deve ser o mesmo que o consumidor no webhook-processor-ms está ouvindo.
	queueName := "whatsapp-webhooks-pf-raw"

	// O publisher injetado é usado para enviar a mensagem.
	err := s.publisher.Publish(ctx, queueName, payload)
	if err != nil {
		slog.Error("Falha ao publicar webhook na fila", "error", err, "queue", queueName)
		return err
	}

	slog.Info("Webhook publicado com sucesso na fila", "queue", queueName)
	return nil
}
