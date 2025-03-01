package migrations

import "gorm.io/gorm"

func AddPhotoIDColumn(db *gorm.DB) error {
	// Check if column exists
	var count int64
	db.Raw(`SELECT COUNT(*) 
            FROM INFORMATION_SCHEMA.COLUMNS 
            WHERE TABLE_SCHEMA = 'evermos_internship' 
            AND TABLE_NAME = 'foto_produk' 
            AND COLUMN_NAME = 'photo_id'`).Count(&count)

	if count == 0 {
		// Column doesn't exist, so add it
		return db.Exec(`
            ALTER TABLE foto_produk 
            ADD COLUMN photo_id BIGINT AFTER id_produk;
        `).Error
	}

	return nil
}
