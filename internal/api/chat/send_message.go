package chat // Определение пакета

import (
	"context"                                          // Импорт контекста для управления жизненным циклом запроса
	desc "github.com/Murat993/chat-server/pkg/chat_v1" // Импорт сгенерированных protobuf структур

	"google.golang.org/grpc/codes"                   // Импорт кодов состояния gRPC для указания результатов операции
	"google.golang.org/grpc/status"                  // Импорт статуса gRPC для создания ответов с ошибками
	"google.golang.org/protobuf/types/known/emptypb" // Импорт пакета для возвращения пустого ответа
)

// SendMessage - метод отправки сообщения в чат
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	i.mxChannel.RLock()                         // Блокировка для чтения из карты каналов сообщений
	chatChan, ok := i.channels[req.GetChatId()] // Получение канала сообщений для указанного чата
	i.mxChannel.RUnlock()                       // Разблокировка после чтения

	if !ok { // Проверка существования чата
		// Возвращение ошибки, если чат не найден
		return nil, status.Errorf(codes.NotFound, "chat not found")
	}

	chatChan <- req.GetMessage() // Отправка сообщения в канал чата

	// Возвращение пустого ответа, указывающего на успешное выполнение метода
	return &emptypb.Empty{}, nil
}
