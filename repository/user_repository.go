package repository

import (
    "ByTeora-Pos-Backend-App/config"
	"ByTeora-Pos-Backend-App/models"
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

func GetUserByEmail(email string) (*models.User, error) {
    var user models.User

    query := `
        SELECT 
            uuid,
            email,
            password,
            nama_depan,
            nama_belakang,
            role,
            status,
            created_at,
            modified_at,
            deleted_at
        FROM user 
        WHERE email = ? AND deleted_at IS NULL
        LIMIT 1
    `

    err := config.DB.QueryRow(query, email).Scan(
        &user.UUID,
        &user.Email,
        &user.Password,
        &user.NamaDepan,
        &user.NamaBelakang,
        &user.Role,
        &user.Status,
        &user.CreatedAt,
        &user.ModifiedAt,
        &user.DeletedAt,
    )

    if err != nil {
        return nil, err
    }

    return &user, nil
}
