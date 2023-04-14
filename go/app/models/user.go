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
