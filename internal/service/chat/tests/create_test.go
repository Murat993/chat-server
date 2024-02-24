package tests

import (
	"context"
	"fmt"
	"github.com/Murat993/chat-server/internal/client/db"
	txMock "github.com/Murat993/chat-server/internal/client/db/mocks"
	"github.com/Murat993/chat-server/internal/dto"
	"github.com/Murat993/chat-server/internal/repository"
	"github.com/Murat993/chat-server/internal/repository/mocks"
	"github.com/Murat993/chat-server/internal/service/chat"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type chatTxManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *dto.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Int64()
		title   = gofakeit.Animal()
		content = gofakeit.Animal()

		repoErr = fmt.Errorf("repo error")

		req = &dto.Chat{
			Usernames: []string{title, content},
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		txManagerMock      chatTxManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)

			service := chat.NewMockService(chatRepoMock, txManagerMock)

			newID, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
