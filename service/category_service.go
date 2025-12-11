package service

import (
	"ByTeora-Pos-Backend-App/config"
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
