package models

import "time"

type Product struct {
	ID          int        `json:"id"`
	UUID        string     `json:"uuid"`
	StoreID     int        `json:"store_id"`
	CategoryID  *int       `json:"category_id,omitempty"`
	ProductName string     `json:"product_name"`
	SKU         *string    `json:"sku,omitempty"`
	Barcode     *string    `json:"barcode,omitempty"`
	Description *string    `json:"description,omitempty"`
	Price       float64    `json:"price"`
	Cost        *float64   `json:"cost,omitempty"`
	Status      string     `json:"status"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	ModifiedAt  time.Time  `json:"modified_at"`
}