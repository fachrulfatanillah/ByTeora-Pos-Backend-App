package service

import (
	"ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/models"
	"ByTeora-Pos-Backend-App/api/request"
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

func IsStoreOwnedByUser(storeUUID, userUUID string) (bool, error) {
    query := `
        SELECT COUNT(*)
        FROM store s
        JOIN user u ON s.user_id = u.id
        WHERE s.uuid = ? AND u.uuid = ? AND s.deleted_at IS NULL
    `

    var count int
    err := config.DB.QueryRow(query, storeUUID, userUUID).Scan(&count)
    return count > 0, err
}

func UpdateStore(storeUUID string, data request.UpdateStoreRequest) error {
    query := `
        UPDATE store
        SET store_name = ?, address = ?, phone_number = ?, status = ?, modified_at = NOW()
        WHERE uuid = ?
    `

    _, err := config.DB.Exec(
        query,
        data.StoreName,
        data.Address,
        data.PhoneNumber,
        data.Status,
        storeUUID,
    )

    return err
}

func GetStoreByUUID(storeUUID string) (models.Store, error) {
    var s models.Store

    query := `
        SELECT uuid, store_name, address, phone_number, status
        FROM store
        WHERE uuid = ? AND deleted_at IS NULL
    `

    err := config.DB.QueryRow(query, storeUUID).Scan(
        &s.UUID,
        &s.StoreName,
        &s.Address,
        &s.PhoneNumber,
        &s.Status,
    )

    return s, err
}

func SoftDeleteStore(storeUUID string) error {
    query := `
        UPDATE store
        SET deleted_at = NOW()
        WHERE uuid = ? AND deleted_at IS NULL
    `

    _, err := config.DB.Exec(query, storeUUID)
    return err
}