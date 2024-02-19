package chat

import (
	"context"
	"github.com/Murat993/chat-server/internal/dto"
)

func (s server) Create(ctx context.Context, chat *dto.Chat) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx, chat)
		if errTx != nil {
			return errTx
		}

		//_, errTx = s.chatRepository.Get(ctx, id)
		//if errTx != nil {
		//	return errTx
		//}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
