package tests

// Импортируем необходимые пакеты: стандартные библиотеки Go, библиотеки для генерации данных, мокирования, утверждений,
// а также внутренние модули для работы с API, моделями и сервисами.
import (
	"context"
	"fmt"
	"github.com/Murat993/chat-server/internal/api/chat"
	"github.com/Murat993/chat-server/internal/dto"
	"github.com/Murat993/chat-server/internal/service"
	"github.com/Murat993/chat-server/internal/service/mocks"
	desc "github.com/Murat993/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"testing"

	"github.com/gojuno/minimock/v3"       // Для мокирования зависимостей
	"github.com/stretchr/testify/require" // Для утверждений в тестах
	// Внутренние пакеты проекта
)

// TestCreate - функция для тестирования создания заметки.
func TestCreate(t *testing.T) {
	t.Parallel() // Запускаем тесты параллельно для ускорения выполнения.

	// Определяем типы и переменные для использования в тестах.
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	//  Структура для аргументов, которые будут переданы в тестируемую функцию.
	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	// Подготавливаем контекст, минимок контроллер и тестовые данные.
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Int64() // Генерируем тестовые данные
		title   = gofakeit.Animal()
		content = gofakeit.Animal()

		serviceErr = fmt.Errorf("service error") // Ошибка для имитации сбоя в сервисе

		// Создаем запрос на создание заметки
		req = &desc.CreateRequest{
			Usernames: []string{title, content},
		}

		// Информация для создания заметки
		chatDto = &dto.Chat{
			Usernames: []string{title, content},
		}

		// Ожидаемый ответ от сервиса
		res = &desc.CreateResponse{
			Id: id,
		}
	)
	defer t.Cleanup(mc.Finish) // Очищаем ресурсы после выполнения каждого теста.

	// Определяем тестовые случаи.
	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		// Тест на успешное создание заметки
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatDto).Return(id, nil) // Настраиваем мок для имитации успешного создания
				return mock
			},
		},
		// Тест на случай ошибки в сервисе
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatDto).Return(0, serviceErr) // Настраиваем мок для имитации ошибки сервиса
				return mock
			},
		},
	}

	// Итерируем по тестовым случаям и запускаем каждый тест.
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Запускаем тесты в параллельном режиме для каждого тестового случая.

			// Создаем мок сервиса заметок с помощью переданной функции.
			chatServiceMock := tt.chatServiceMock(mc)
			// Инициализируем API с моком сервиса.
			api := chat.NewImplementation(chatServiceMock)

			// Вызываем метод Create API и получаем результат.
			newID, err := api.Create(tt.args.ctx, tt.args.req)

			// Проверяем, соответствует ли полученная ошибка ожидаемой ошибке.
			require.Equal(t, tt.err, err)
			// Проверяем, соответствует ли возвращенный результат ожидаемому результату.
			require.Equal(t, tt.want, newID)
		})
	}
}
