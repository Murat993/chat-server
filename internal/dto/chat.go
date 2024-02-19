package dto

import (
	"time"
)

type Chat struct {
	Usernames []string
}

type Message struct {
	From      string
	Text      string
	CreatedAt time.Time
}
