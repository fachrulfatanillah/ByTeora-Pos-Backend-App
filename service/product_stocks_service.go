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

func GetAllProductStockLogs(storeUUID string) ([]response.ProductStockLogResponse, error) {
    query := `
        SELECT 
            ps.uuid,
            p.uuid,
            s.uuid,
            ps.stock_in,
            ps.stock_out,
            ps.status,
            ps.created_at
        FROM product_stocks ps
        JOIN product p ON ps.product_id = p.id
        JOIN store s ON ps.store_id = s.id
        WHERE s.uuid = ?
        ORDER BY ps.created_at DESC
    `

    rows, err := config.DB.Query(query, storeUUID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []response.ProductStockLogResponse

    for rows.Next() {
        var item response.ProductStockLogResponse

        err := rows.Scan(
            &item.StockUUID,
            &item.ProductUUID,
            &item.StoreUUID,
            &item.StockIn,
            &item.StockOut,
            &item.Status,
            &item.CreatedAt,
        )
        if err != nil {
            return nil, err
        }

        results = append(results, item)
    }

    return results, nil
}

func GetCurrentStockByUUID(productUUID, storeUUID string) (int, error) {
    query := `
        SELECT 
            SUM(stock_in) - SUM(stock_out) AS current_stock
        FROM product_stocks ps
        JOIN product p ON ps.product_id = p.id
        JOIN store s ON ps.store_id = s.id
        WHERE p.uuid = ? AND s.uuid = ?
    `

    var currentStock int
    err := config.DB.QueryRow(query, productUUID, storeUUID).Scan(&currentStock)
    if err != nil {
        return 0, err
    }

    return currentStock, nil
}

func CheckStoreOwnership(userID int, storeUUID string) (bool, error) {
    var count int

    query := `
        SELECT COUNT(*)
        FROM store
        WHERE uuid = ? AND user_id = ? AND deleted_at IS NULL
    `

    err := config.DB.QueryRow(query, storeUUID, userID).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func GetProductStockLogsByProduct(storeUUID, productUUID string) ([]response.ProductStockLogResponse, error) {
    query := `
        SELECT 
            ps.uuid,
            p.uuid,
            s.uuid,
            ps.stock_in,
            ps.stock_out,
            ps.status,
            ps.created_at
        FROM product_stocks ps
        JOIN product p ON ps.product_id = p.id
        JOIN store s ON ps.store_id = s.id
        WHERE s.uuid = ? AND p.uuid = ?
        ORDER BY ps.created_at DESC
    `

    rows, err := config.DB.Query(query, storeUUID, productUUID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []response.ProductStockLogResponse

    for rows.Next() {
        var item response.ProductStockLogResponse
        err := rows.Scan(
            &item.StockUUID,
            &item.ProductUUID,
            &item.StoreUUID,
            &item.StockIn,
            &item.StockOut,
            &item.Status,
            &item.CreatedAt,
        )
        if err != nil {
            return nil, err
        }

        results = append(results, item)
    }

    return results, nil
}

func GetAllProductStocks(storeUUID string) ([]response.ProductStockCurrentResponse, error) {
	query := `
		SELECT 
			p.uuid AS product_uuid,
			s.uuid AS store_uuid,
			COALESCE(SUM(ps.stock_in), 0) - COALESCE(SUM(ps.stock_out), 0) AS current_stock,
			COALESCE(MAX(ps.uuid), '') AS stock_uuid
		FROM product p
		JOIN store s ON p.store_id = s.id
		LEFT JOIN product_stocks ps ON ps.product_id = p.id AND ps.deleted_at IS NULL
		WHERE s.uuid = ?
		  AND p.deleted_at IS NULL
		GROUP BY p.id, s.id
		ORDER BY p.created_at DESC
	`

	rows, err := config.DB.Query(query, storeUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []response.ProductStockCurrentResponse

	for rows.Next() {
		var item response.ProductStockCurrentResponse
		err := rows.Scan(&item.ProductUUID, &item.StoreUUID, &item.CurrentStock, &item.StockUUID)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, nil
}

