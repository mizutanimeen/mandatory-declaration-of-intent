package models

import (
	"time"
)

type User struct {
	UserID   int       `json:"userid"`
	Name     string    `json:"name"`
	Text     string    `json:"text"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type GestUser struct {
	GestUserID int       `json:"gest_userid"`
	Name       string    `json:"name"`
	Text       string    `json:"text"`
	RoomID     string    `json:"roomid"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
