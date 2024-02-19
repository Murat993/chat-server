package repository

import (
	"context"
	"github.com/Murat993/chat-server/internal/dto"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *dto.Chat) (int64, error)
}
