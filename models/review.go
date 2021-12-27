package models

import "time"

type (
	Review struct {
		ID          uint      `json:"id" gorm:"primary_key"`
		UserID      uint      `json:"user_id" gorm:"unique"`
		GameID      uint      `json:"game_id" gorm:"unique"`
		Rating      int       `json:"rating"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		User        Company   `json:"-"`
		Game        Company   `json:"-"`
	}
)
