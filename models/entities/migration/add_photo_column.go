package migration

import (
	"fmt"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

func AddPhotoColumnToFotoToko(db *gorm.DB) error {
	// Debug: Print current schema
	var columns []map[string]interface{}
	db.Raw("SHOW COLUMNS FROM foto_toko").Scan(&columns)
	fmt.Println("Current foto_toko table schema:")
	for _, col := range columns {
		fmt.Printf("%v\n", col)
	}

	// Check and fix id_toko column if needed
	if !db.Migrator().HasColumn(&entities.StorePhoto{}, "id_toko") {
		fmt.Println("Adding id_toko column...")
		db.Exec("ALTER TABLE foto_toko ADD COLUMN id_toko BIGINT UNSIGNED NOT NULL")
		db.Exec("ALTER TABLE foto_toko ADD INDEX idx_foto_toko_id_toko (id_toko)")
	}

	// Add photo column if it doesn't exist
	if !db.Migrator().HasColumn(&entities.StorePhoto{}, "photo") {
		fmt.Println("Adding photo column...")
		return db.Exec("ALTER TABLE foto_toko ADD COLUMN photo VARCHAR(255)").Error
	}

	return nil
}
