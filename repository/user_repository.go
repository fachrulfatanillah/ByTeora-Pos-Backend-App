package repository

import (
    "ByTeora-Pos-Backend-App/config"
)

func CountUsers() (int, error) {
    var total int
    err := config.DB.QueryRow("SELECT COUNT(*) FROM user WHERE deleted_at IS NULL").Scan(&total)
    return total, err
}

func IsEmailExists(email string) (bool, error) {
    var count int
    err := config.DB.QueryRow(
        "SELECT COUNT(*) FROM user WHERE email = ? AND deleted_at IS NULL",
        email,
    ).Scan(&count)

    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func CreateUser(
    uuid string,
    email string,
    passwordHash string,
    namaDepan string,
    namaBelakang string,
    role string,
) error {
    _, err := config.DB.Exec(`
        INSERT INTO user (uuid, email, password, nama_depan, nama_belakang, role, status, created_at, modified_at)
        VALUES (?, ?, ?, ?, ?, ?, 'active', NOW(), NOW())
    `,
        uuid, email, passwordHash, namaDepan, namaBelakang, role,
    )

    return err
}