package migration

import (
	"log"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	// Create tables
	tables := []interface{}{
		&entities.User{},
		&entities.Address{},
		&entities.Category{},
		&entities.Store{},
		&entities.StorePhoto{},
		&entities.Product{},
		&entities.FotoProduk{},
		&entities.Trx{},
		&entities.TrxDetail{},
		&entities.ProductLog{},
		&entities.KeranjangBelanja{},
		&entities.Wishlist{},
		&entities.ProductReview{},
		&entities.Notification{},
		&entities.ProductPromo{},
		&entities.DiskonProduk{},
		&entities.Order{},
		&entities.ProductCoupon{},
	}

	// Run migrations for all tables
	for _, table := range tables {
		err := db.AutoMigrate(table)
		if err != nil {
			log.Fatalf("Failed to migrate table: %v", err)
		}
	}

	// Create log_produk table with custom SQL
	err := createLogProdukTable(db)
	if err != nil {
		log.Fatalf("Failed to create log_produk table: %v", err)
	}
}
