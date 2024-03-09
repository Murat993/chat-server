package chat // Определение пакета

import (
	"context"                                          // Импорт пакета для контекста
	desc "github.com/Murat993/chat-server/pkg/chat_v1" // Импорт сгенерированных protobuf структур

	"github.com/google/uuid"                         // Импорт пакета для генерации UUID
	"google.golang.org/protobuf/types/known/emptypb" // Импорт пакета для работы с пустым protobuf сообщением
)

// CreateChat - метод создания нового чата
func (i *Implementation) Create(ctx context.Context, _ *emptypb.Empty) (*desc.CreateResponse, error) {
	chatID, err := uuid.NewUUID() // Генерация нового UUID для идентификатора чата
	if err != nil {
		return nil, err // Возвращение ошибки, если не удалось сгенерировать UUID
	}

	i.channels[chatID.String()] = make(chan *desc.Message, 100) // Создание канала сообщений для чата с буфером на 100 сообщений

	// Возвращение ответа с идентификатором созданного чата
	return &desc.CreateResponse{
		ChatId: chatID.String(), // Установка идентификатора чата в ответ
	}, nil
}

// chat (chat_id)(u1, u2, u3...)
// u1 -> chat stream
// u2 -> chat stream
// u3 -> chat stream
