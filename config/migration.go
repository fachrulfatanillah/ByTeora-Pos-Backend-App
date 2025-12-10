package config

import (
	"log"
)

func RunMigrations() {
	userTable := `
	CREATE TABLE IF NOT EXISTS user (
		id INT AUTO_INCREMENT PRIMARY KEY,
		uuid VARCHAR(36) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		nama_depan VARCHAR(100),
		nama_belakang VARCHAR(100),
		image_url TEXT,
		role ENUM('admin', 'cashier', 'owner') NOT NULL,
		status ENUM('active', 'inactive', 'banned') NOT NULL DEFAULT 'active',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
	`

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration completed!")
}
