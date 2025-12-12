package service

import (
	"ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/api/request"
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