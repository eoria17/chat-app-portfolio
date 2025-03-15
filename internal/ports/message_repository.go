package ports

import "github.com/eoria17/chat-app-portfolio/internal/core/domain"

type MessageRepository interface {
	SaveMessage(msg domain.Message) error
	GetMessages() ([]domain.Message, error)
}
