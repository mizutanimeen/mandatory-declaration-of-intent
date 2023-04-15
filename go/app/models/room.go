package models

import (
	"time"
)

type Room struct {
	RoomID      string    `json:"roomid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CookieName  string    `json:"cookie_name"`
	CookieValue string    `json:"cookie_value"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
