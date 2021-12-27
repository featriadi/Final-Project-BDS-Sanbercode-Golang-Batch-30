package models

import "time"

type (
	Game struct {
		ID          uint        `json:"id" gorm:"primary_key"`
		Title       string      `json:"title"`
		DeveloperID uint        `json:"developer_id"`
		PublisherID uint        `json:"published_id"`
		Price       int         `json:"price"`
		Year        int         `json:"year"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		Developer   Company     `json:"-"`
		Publisher   Company     `json:"-"`
		GameImages  []GameImage `json:"-"`
		GameTags    []GameTag   `json:"-"`
		Reviews     []Review    `json:"-"`
	}
)
