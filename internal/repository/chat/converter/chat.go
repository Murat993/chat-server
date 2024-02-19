package converter

import (
	"github.com/Murat993/chat-server/internal/dto"
	modelRepo "github.com/Murat993/chat-server/internal/repository/chat/entity"
)

func ToChatFromRepo(chat *modelRepo.Chat) *dto.Chat {
	return &dto.Chat{
		Usernames: chat.Usernames,
	}
}
