package chat

import (
	"github.com/Murat993/chat-server/internal/client/db"
	"github.com/Murat993/chat-server/internal/repository"
	def "github.com/Murat993/chat-server/internal/service"
)

var _ def.ChatService = (*server)(nil)

type server struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(chatRepository repository.ChatRepository, txManager db.TxManager) *server {
	return &server{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}

func NewMockService(deps ...interface{}) def.ChatService {
	srv := server{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.ChatRepository:
			srv.chatRepository = s
		case db.TxManager:
			srv.txManager = s
		}
	}

	return &srv
}
