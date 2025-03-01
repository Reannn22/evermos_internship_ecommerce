package database

import (
	"mini-project-evermos/database/migrations"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	// ...existing migrations...
	migrations.CreateKeranjangBelanjaTable(db)
	migrations.CreateProductReviewsTable(db)
}
