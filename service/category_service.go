package service

import (
	"ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/models"
	"ByTeora-Pos-Backend-App/api/request"
)

func GetStoreIDByUUID(storeUUID string) (int, error) {
	var id int
	err := config.DB.QueryRow(`SELECT id FROM store WHERE uuid = ? AND deleted_at IS NULL`, storeUUID).Scan(&id)
	return id, err
}

func CreateCategory(categoryUUID string, storeID int, name, description string) error {
	_, err := config.DB.Exec(`
		INSERT INTO category (uuid, store_id, category_name, description, status, created_at, modified_at)
		VALUES (?, ?, ?, ?, 'active', NOW(), NOW())
	`, categoryUUID, storeID, name, description)
	return err
}

func GetCategoriesByStoreID(storeID int) ([]models.Category, error) {
	query := `
		SELECT uuid, category_name, description, status, created_at, modified_at
		FROM category
		WHERE store_id = ? AND deleted_at IS NULL
	`
	rows, err := config.DB.Query(query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		err := rows.Scan(
			&c.UUID,
			&c.CategoryName,
			&c.Description,
			&c.Status,
			&c.CreatedAt,
			&c.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func GetCategoryByUUID(categoryUUID string) (models.Category, error) {
	var cat models.Category

	query := `
		SELECT c.id, c.uuid, c.store_id, c.category_name, c.description, c.status, c.deleted_at, c.created_at, c.modified_at
		FROM category c
		WHERE c.uuid = ? AND c.deleted_at IS NULL
	`

	err := config.DB.QueryRow(query, categoryUUID).Scan(
		&cat.ID,
		&cat.UUID,
		&cat.StoreID,
		&cat.CategoryName,
		&cat.Description,
		&cat.Status,
		&cat.DeletedAt,
		&cat.CreatedAt,
		&cat.ModifiedAt,
	)

	return cat, err
}

func UpdateCategory(categoryUUID string, data request.UpdateCategoryRequest) error {
	query := `
		UPDATE category
		SET category_name = ?, description = ?, status = ?, modified_at = NOW()
		WHERE uuid = ? AND deleted_at IS NULL
	`

	_, err := config.DB.Exec(
		query,
		data.CategoryName,
		data.Description,
		data.Status,
		categoryUUID,
	)

	return err
}