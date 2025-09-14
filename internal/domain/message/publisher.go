package message

import "context"

// MessagePublisher define a interface para publicar mensagens.
type MessagePublisher interface {
	Publish(context.Context, string, []byte) error
	Close() error
}
