package models

import "time"

type (
	GameTag struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		GameID    uint      `json:"game_id"`
		TagID     uint      `json:"tag_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Game      Company   `json:"-"`
		Tag       Company   `json:"-"`
	}
)
