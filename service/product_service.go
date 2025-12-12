package service

import (
	"ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/api/request"
	"ByTeora-Pos-Backend-App/api/response"
)

func CreateProduct(productUUID string, storeID int, categoryID *int, req request.CreateProductRequest, status string) error {
	_, err := config.DB.Exec(`
		INSERT INTO product (uuid, store_id, category_id, product_name, sku, barcode, description, price, cost, status, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`, productUUID, storeID, categoryID, req.ProductName, req.SKU, req.Barcode, req.Description, req.Price, req.Cost, status)

	return err
}

func GetCategoryIDByUUID(categoryUUID string) (int, error) {
	var id int
	err := config.DB.QueryRow(`SELECT id FROM category WHERE uuid = ? AND deleted_at IS NULL`, categoryUUID).Scan(&id)
	return id, err
}

func GetProductsByStoreUUID(storeUUID string) ([]response.ProductItemResponse, error) {

	rows, err := config.DB.Query(`
		SELECT 
			p.id,
			p.uuid AS product_uuid,
			s.uuid AS store_uuid,
			c.uuid AS category_uuid,
			p.product_name,
			p.sku,
			p.barcode,
			p.description,
			p.price,
			p.cost,
			p.status,
			DATE_FORMAT(p.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
			DATE_FORMAT(p.modified_at, '%Y-%m-%d %H:%i:%s') AS updated_at
		FROM product p
		JOIN store s ON s.id = p.store_id
		LEFT JOIN category c ON c.id = p.category_id
		WHERE s.uuid = ? AND p.deleted_at IS NULL
		ORDER BY p.created_at DESC
	`, storeUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []response.ProductItemResponse

	for rows.Next() {
		var (
			productID int
			product   response.ProductItemResponse
		)

		err := rows.Scan(
			&productID,
			&product.ProductUUID,
			&product.StoreUUID,
			&product.CategoryUUID,
			&product.ProductName,
			&product.SKU,
			&product.Barcode,
			&product.Description,
			&product.Price,
			&product.Cost,
			&product.Status,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		stock, err := GetStockByProductID(productID)
		if err != nil {
			return nil, err
		}
		product.Stock = stock

		products = append(products, product)
	}

	return products, nil
}

func GetProductByUUID(productUUID string, storeID int) (*response.UpdateProductResponse, error) {
	query := `
		SELECT 
			p.uuid,
			s.uuid,
			c.uuid,
			p.product_name,
			p.sku,
			p.barcode,
			p.description,
			p.price,
			p.cost,
			p.status,
			DATE_FORMAT(p.modified_at, '%Y-%m-%d %H:%i:%s')
		FROM product p
		JOIN store s ON s.id = p.store_id
		LEFT JOIN category c ON c.id = p.category_id
		WHERE p.uuid = ? AND p.store_id = ? AND p.deleted_at IS NULL
	`

	row := config.DB.QueryRow(query, productUUID, storeID)

	var res response.UpdateProductResponse
	err := row.Scan(
		&res.ProductUUID,
		&res.StoreUUID,
		&res.CategoryUUID,
		&res.ProductName,
		&res.SKU,
		&res.Barcode,
		&res.Description,
		&res.Price,
		&res.Cost,
		&res.Status,
		&res.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func UpdateProductPartial(productUUID string, storeID int, req request.UpdateProductRequest) error {
	query := "UPDATE product SET "
	args := []interface{}{}
	cols := []string{}

	if req.CategoryUUID != nil {
		var categoryID int
		_ = config.DB.QueryRow("SELECT id FROM category WHERE uuid = ?", *req.CategoryUUID).Scan(&categoryID)

		cols = append(cols, "category_id = ?")
		args = append(args, categoryID)
	}

	if req.ProductName != nil {
		cols = append(cols, "product_name = ?")
		args = append(args, *req.ProductName)
	}
	if req.SKU != nil {
		cols = append(cols, "sku = ?")
		args = append(args, *req.SKU)
	}
	if req.Barcode != nil {
		cols = append(cols, "barcode = ?")
		args = append(args, *req.Barcode)
	}
	if req.Description != nil {
		cols = append(cols, "description = ?")
		args = append(args, *req.Description)
	}
	if req.Price != nil {
		cols = append(cols, "price = ?")
		args = append(args, *req.Price)
	}
	if req.Cost != nil {
		cols = append(cols, "cost = ?")
		args = append(args, *req.Cost)
	}
	if req.Status != nil {
		cols = append(cols, "status = ?")
		args = append(args, *req.Status)
	}

	if len(cols) == 0 {
		return nil
	}

	cols = append(cols, "modified_at = NOW()")

	for i, col := range cols {
		if i == 0 {
			query += col
		} else {
			query += ", " + col
		}
	}

	query += " WHERE uuid = ? AND store_id = ? AND deleted_at IS NULL"

	args = append(args, productUUID, storeID)

	_, err := config.DB.Exec(query, args...)
	return err
}

func SoftDeleteProduct(productUUID string, storeID int) error {
	query := `
		UPDATE product 
		SET deleted_at = NOW() 
		WHERE uuid = ? AND store_id = ? AND deleted_at IS NULL
	`

	_, err := config.DB.Exec(query, productUUID, storeID)
	return err
}

func IsProductBelongsToStore(productUUID string, storeID int) (bool, error) {
	var count int
	err := config.DB.QueryRow(`
		SELECT COUNT(*) FROM product 
		WHERE uuid = ? AND store_id = ? AND deleted_at IS NULL
	`, productUUID, storeID).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetStockByProductID(productID int) (int, error) {
    var stock int

    err := config.DB.QueryRow(`
        SELECT COALESCE(SUM(stock_in) - SUM(stock_out), 0)
        FROM product_stocks
        WHERE product_id = ? AND deleted_at IS NULL
    `, productID).Scan(&stock)

    if err != nil {
        return 0, err
    }

    return stock, nil
}