package migrations

import (
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

func CreateKeranjangBelanjaTable(db *gorm.DB) error {
	return db.AutoMigrate(&entities.KeranjangBelanja{})
}
