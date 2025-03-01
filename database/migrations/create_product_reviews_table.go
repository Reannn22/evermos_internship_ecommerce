package migrations

import (
	"gorm.io/gorm"
)

func CreateProductReviewsTable(db *gorm.DB) error {
	sql := `
	DROP TABLE IF EXISTS product_reviews;
	CREATE TABLE product_reviews (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		id_toko BIGINT UNSIGNED NOT NULL,
		id_produk INT UNSIGNED NOT NULL,
		ulasan TEXT NOT NULL,
		rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT fk_product_reviews_store FOREIGN KEY (id_toko) REFERENCES toko(id) ON DELETE CASCADE,
		CONSTRAINT fk_product_reviews_product FOREIGN KEY (id_produk) REFERENCES produk(id) ON DELETE CASCADE
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`
	return db.Exec(sql).Error
}
