package models

import "time"

type (
	GameImage struct {
		ID        uint      `json:"id"`
		GameID    uint      `json:"game_id"`
		ImageUrl  string    `json:"image_url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Game      Company   `json:"-"`
	}
)
