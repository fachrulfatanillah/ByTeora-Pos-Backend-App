package repository

import (
	"ByTeora-Pos-Backend-App/config"
)

func GetUserIDByUUID(userUUID string) (int, error) {
	var id int
	err := config.DB.QueryRow(`
		SELECT id FROM user WHERE uuid = ? AND deleted_at IS NULL
	`, userUUID).Scan(&id)
	return id, err
}

func CreateStore(storeUUID string, userID int, storeName, address, phone string) error {
	_, err := config.DB.Exec(`
		INSERT INTO store (uuid, user_id, store_name, address, phone_number, status, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, 'active', NOW(), NOW())
	`,
		storeUUID, userID, storeName, address, phone,
	)

	return err
}