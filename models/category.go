package models

import "time"

type Category struct {
	ID           int       `json:"id"`
	UUID         string    `json:"uuid"`
	StoreID      int       `json:"store_id"`
	CategoryName string    `json:"category_name"`
	Description  string    `json:"description,omitempty"`
	Status       string    `json:"status"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	ModifiedAt   time.Time  `json:"modified_at"`
}