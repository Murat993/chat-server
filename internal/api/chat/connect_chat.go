package chat

import (
	desc "github.com/Murat993/chat-server/pkg/chat_v1" // Сгенерированные protobuf-структуры
	"google.golang.org/grpc/codes"                     // Импорт для использования стандартных кодов ошибок gRPC
	"google.golang.org/grpc/status"                    // Импорт для создания статусов ошибок gRPC
)

// ConnectChat обрабатывает запрос на подключение к чату
func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	// Поиск канала сообщений чата по идентификатору
	i.mxChannel.RLock()                         // Блокировка для чтения из карты каналов сообщений
	chatChan, ok := i.channels[req.GetChatId()] // Получение канала сообщений для указанного чата
	i.mxChannel.RUnlock()                       // Разблокировка после чтения

	if !ok {
		// Возврат ошибки, если чат не найден
		return status.Errorf(codes.NotFound, "chat not found")
	}

	// Инициализация структуры чата, если она еще не была создана
	i.mxChat.Lock()
	if _, okChat := i.chats[req.GetChatId()]; !okChat {
		i.chats[req.GetChatId()] = &Chat{
			streams: make(map[string]desc.ChatV1_ConnectChatServer),
		}
	}
	i.mxChat.Unlock()

	// Регистрация потока пользователя в чате
	i.chats[req.GetChatId()].m.Lock()
	i.chats[req.GetChatId()].streams[req.GetUsername()] = stream
	i.chats[req.GetChatId()].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChan:
			// Чтение сообщений из канала чата
			if !okCh {
				// Выход из цикла, если канал закрыт
				return nil
			}

			// Рассылка сообщения всем подключенным пользователям чата
			for _, st := range i.chats[req.GetChatId()].streams {
				if err := st.Send(msg); err != nil {
					// Возврат ошибки в случае неудачной отправки сообщения
					return err
				}
			}

		case <-stream.Context().Done():
			// Удаление потока пользователя при отключении
			i.chats[req.GetChatId()].m.Lock()
			delete(i.chats[req.GetChatId()].streams, req.GetUsername())
			i.chats[req.GetChatId()].m.Unlock()
			return nil
		}
	}
}
