package converter

import (
	"github.com/Murat993/chat-server/internal/dto"
)

//func ToChatFromService(note *entity.Note) *desc.Note {
//	Chat
//	return &desc.Note{
//		Id:        note.ID,
//		CreatedAt: timestamppb.New(note.CreatedAt),
//		UpdatedAt: updatedAt,
//	}
//}

func ToChatFromDesc(usernames []string) *dto.Chat {
	return &dto.Chat{
		Usernames: usernames,
	}
}
