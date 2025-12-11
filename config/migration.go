package config

import (
	"log"
)

func RunMigrations() {
	MigrateUserTable()
	MigrateStoreTable()
	MigrateCategoryTable()
}

func MigrateUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS user (
		id INT AUTO_INCREMENT PRIMARY KEY,
		uuid VARCHAR(36) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		nama_depan VARCHAR(100),
		nama_belakang VARCHAR(100),
		image_url TEXT,
		role ENUM('admin', 'owner') NOT NULL,
		status ENUM('active', 'inactive', 'banned') NOT NULL DEFAULT 'active',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration completed!")
}

func MigrateStoreTable() {
	query := `
	CREATE TABLE IF NOT EXISTS store (
		id INT AUTO_INCREMENT PRIMARY KEY,
		uuid VARCHAR(36) NOT NULL UNIQUE,
		user_id INT NOT NULL,
		store_name VARCHAR(150) NOT NULL,
		address TEXT NULL,
		phone_number VARCHAR(20) NULL,
		status VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
		deleted_at TIMESTAMP NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		
		CONSTRAINT fk_store_user FOREIGN KEY (user_id) REFERENCES user(id)
	) ENGINE=InnoDB;
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate store table:", err)
	}

	log.Println("Store table migrated successfully!")
}

func MigrateCategoryTable() {
	query := `
		CREATE TABLE IF NOT EXISTS category (
		id INT AUTO_INCREMENT PRIMARY KEY,
		uuid VARCHAR(36) NOT NULL UNIQUE,
		store_id INT NOT NULL,
		category_name VARCHAR(150) NOT NULL,
		description TEXT NULL,
		status ENUM('active','inactive') NOT NULL DEFAULT 'active',
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT fk_category_store FOREIGN KEY (store_id) REFERENCES store(id)
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate store table:", err)
	}

	log.Println("Store table migrated successfully!")
}