package migrations

import "gorm.io/gorm"

func CreateFotoProdukTable(db *gorm.DB) error {
	sql := `
    CREATE TABLE IF NOT EXISTS foto_produk (
        id BIGINT PRIMARY KEY AUTO_INCREMENT,
        id_produk BIGINT NOT NULL,
        photo_id BIGINT,
        url VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY (id_produk) REFERENCES produk(id)
    );
    `
	return db.Exec(sql).Error
}
