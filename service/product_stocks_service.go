package service

import (
	"ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/api/request"
	"ByTeora-Pos-Backend-App/api/response"

	"github.com/google/uuid"
)

func CreateProductStock(storeID, productID int, req request.CreateProductStockRequest) (*response.ProductStockResponse, error) {

	stockIn := 0
	stockOut := 0

	if req.StockIn != nil {
		stockIn = *req.StockIn
	}
	if req.StockOut != nil {
		stockOut = *req.StockOut
	}

	status := "active"

	stockUUID := uuid.New().String()

	query := `
		INSERT INTO product_stocks 
		(uuid, product_id, store_id, stock_in, stock_out, status)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := config.DB.Exec(
		query,
		stockUUID,
		productID,
		storeID,
		stockIn,
		stockOut,
		status,
	)
	if err != nil {
		return nil, err
	}

	currentStock, err := GetCurrentStock(productID, storeID)
	if err != nil {
		return nil, err
	}

	return &response.ProductStockResponse{
		StockUUID:    stockUUID,
		ProductUUID:  "",
		StoreUUID:    "",
		StockIn:      stockIn,
		StockOut:     stockOut,
		CurrentStock: currentStock,
		Status:       status,
	}, nil
}

func GetProductIDByUUID(productUUID string, storeID int) (int, error) {
	var id int

	query := `
		SELECT id 
		FROM product 
		WHERE uuid = ? 
		  AND store_id = ? 
		  AND deleted_at IS NULL
	`

	err := config.DB.QueryRow(query, productUUID, storeID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetCurrentStock(productID, storeID int) (int, error) {
	var totalIn, totalOut int

	query := `
		SELECT 
			COALESCE(SUM(stock_in), 0) AS total_in,
			COALESCE(SUM(stock_out), 0) AS total_out
		FROM product_stocks
		WHERE product_id = ?
		  AND store_id = ?
		  AND deleted_at IS NULL
	`

	err := config.DB.QueryRow(query, productID, storeID).Scan(&totalIn, &totalOut)
	if err != nil {
		return 0, err
	}

	return totalIn - totalOut, nil
}