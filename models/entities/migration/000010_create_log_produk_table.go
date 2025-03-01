package migration

import (
	"fmt"

	"gorm.io/gorm"
)

func createLogProdukTable(db *gorm.DB) error {
	sql := `
	CREATE TABLE IF NOT EXISTS log_produk (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		id_produk BIGINT UNSIGNED NOT NULL,
		nama_produk VARCHAR(255) NOT NULL,
		slug VARCHAR(255) NOT NULL,
		harga_reseller VARCHAR(50) NOT NULL,
		harga_konsumen VARCHAR(50) NOT NULL,
		deskripsi TEXT,
		id_toko BIGINT UNSIGNED NOT NULL,
		id_category BIGINT UNSIGNED NOT NULL,
		kuantitas INT NOT NULL DEFAULT 0,
		harga_total DECIMAL(15,2) NOT NULL DEFAULT 0,
		created_at TIMESTAMP NULL,
		updated_at TIMESTAMP NULL,
		FOREIGN KEY (id_produk) REFERENCES produk(id) ON DELETE CASCADE,
		FOREIGN KEY (id_toko) REFERENCES toko(id) ON DELETE CASCADE,
		FOREIGN KEY (id_category) REFERENCES category(id) ON DELETE CASCADE
	)`

	if err := db.Exec(sql).Error; err != nil {
		return fmt.Errorf("error creating log_produk table: %w", err)
	}
	return nil
}
