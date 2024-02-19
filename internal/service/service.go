package service

import (
	"context"
	"github.com/Murat993/chat-server/internal/dto"
)

type ChatService interface {
	Create(ctx context.Context, info *dto.Chat) (int64, error)
}
