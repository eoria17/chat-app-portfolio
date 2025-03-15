package ports

import "github.com/eoria17/chat-app-portfolio/internal/core/domain"

type AuthService interface {
	LoginGoogle(code string) (*domain.User, error)
}
