package model

import "time"

type ProductStock struct {
	ID           int        `db:"id"`
	UUID         string     `db:"uuid"`
	ProductID    int        `db:"product_id"`
	StoreID      int        `db:"store_id"`
	StockIn      int        `db:"stock_in"`
	StockOut     int        `db:"stock_out"`
	CurrentStock int        `db:"current_stock"`
	Status       string     `db:"status"`
	CreatedAt    time.Time  `db:"created_at"`
	ModifiedAt   time.Time  `db:"modified_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}