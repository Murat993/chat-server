package chat

import (
	desc "github.com/Murat993/chat-server/pkg/chat_v1"
	"sync"
)

// Chat представляет собой структуру чата с потоками сообщений и мьютексом для синхронизации
type Chat struct {
	streams map[string]desc.ChatV1_ConnectChatServer // Карта потоков для каждого пользователя чата
	m       sync.RWMutex                             // Мьютекс для синхронизированного доступа к картам
}

// Implementation реализует интерфейс сервера чата, определенный в protobuf
type Implementation struct {
	desc.UnimplementedChatV1Server // Встраивание для получения реализации по умолчанию

	chats  map[string]*Chat // Карта всех чатов
	mxChat sync.RWMutex     // Мьютекс для синхронизированного доступа к картам чатов

	channels  map[string]chan *desc.Message // Карта каналов сообщений для каждого чата
	mxChannel sync.RWMutex                  // Мьютекс для синхронизированного доступа к каналам сообщений
}

// NewImplementation создает и возвращает новый экземпляр Implementation
func NewImplementation() *Implementation {
	return &Implementation{
		chats:    make(map[string]*Chat),              // Инициализация карты чатов
		channels: make(map[string]chan *desc.Message), // Инициализация карты каналов сообщений
	}
}
