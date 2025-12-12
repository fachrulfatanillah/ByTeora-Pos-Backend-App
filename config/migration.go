package config

import (
	"log"
)

func RunMigrations() {
	MigrateUserTable()
	MigrateStoreTable()
	MigrateCategoryTable()
	MigrateProductTable()
	MigrateProductStockTable()
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

func MigrateProductTable() {
	query := `
		CREATE TABLE IF NOT EXISTS product (
			id INT AUTO_INCREMENT PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL UNIQUE,
			store_id INT NOT NULL,
			category_id INT NULL,
			product_name VARCHAR(200) NOT NULL,
			sku VARCHAR(100) NULL UNIQUE,
			barcode VARCHAR(100) NULL UNIQUE,
			description TEXT NULL,
			price DECIMAL(10,2) NOT NULL DEFAULT 0.00,
			cost DECIMAL(10,2) NULL,
			status ENUM('active', 'inactive') NOT NULL DEFAULT 'active',
			deleted_at TIMESTAMP NULL DEFAULT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			FOREIGN KEY (store_id) REFERENCES store(id),
			FOREIGN KEY (category_id) REFERENCES category(id)
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate store table:", err)
	}

	log.Println("Store table migrated successfully!")
}

func MigrateProductStockTable() {
	query := `
		CREATE TABLE IF NOT EXISTS product_stocks (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL UNIQUE,

			product_id INT NOT NULL,
			store_id INT NOT NULL,

			stock_in INT NOT NULL DEFAULT 0,
			stock_out INT NOT NULL DEFAULT 0,
			current_stock INT NOT NULL DEFAULT 0,

			status VARCHAR(20) NOT NULL DEFAULT 'active',

			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP NULL,

			CONSTRAINT fk_product_stocks_product_id
				FOREIGN KEY (product_id) REFERENCES product(id),

			CONSTRAINT fk_product_stocks_store_id
				FOREIGN KEY (store_id) REFERENCES store(id)
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate store table:", err)
	}

	log.Println("Store table migrated successfully!")
}