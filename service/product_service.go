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
		var product response.ProductItemResponse
		err := rows.Scan(
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

		products = append(products, product)
	}

	return products, nil
}