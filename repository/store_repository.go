package repository

import (
	"ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/models"
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

func GetStoresByUserUUID(userUUID string) ([]models.Store, error) {
	query := `
		SELECT s.uuid, s.store_name, s.address, s.phone_number, s.status
		FROM store s
		JOIN user u ON s.user_id = u.id
		WHERE u.uuid = ? AND s.deleted_at IS NULL
	`

	rows, err := config.DB.Query(query, userUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []models.Store

	for rows.Next() {
		var store models.Store
		err := rows.Scan(
			&store.UUID,
			&store.StoreName,
			&store.Address,
			&store.PhoneNumber,
			&store.Status,
		)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}