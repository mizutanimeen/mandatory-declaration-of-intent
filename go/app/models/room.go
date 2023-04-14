package models

import (
	"time"
)

type Room struct {
	RoomID      int       `json:"roomid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
