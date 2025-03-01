package repositories

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"time"

	"gorm.io/gorm"
)

type ProductPromoRepository interface {
	FindAll() ([]entities.ProductPromo, error)
	FindById(id uint64) (entities.ProductPromo, error)
	FindByProductId(productId uint32) ([]entities.ProductPromo, error)
	Insert(promo models.ProductPromoRequest, storeId uint64) (entities.ProductPromo, error)
	Update(promo models.ProductPromoRequest, id uint64) (entities.ProductPromo, error)
	Delete(id uint64) error
	DeleteAll() error
}

type productPromoRepositoryImpl struct {
	db *gorm.DB
}

func NewProductPromoRepository(db *gorm.DB) ProductPromoRepository {
	return &productPromoRepositoryImpl{db}
}

func (r *productPromoRepositoryImpl) FindAll() ([]entities.ProductPromo, error) {
	var promos []entities.ProductPromo
	err := r.db.Order("id desc").Preload("Store").Preload("Product").Find(&promos).Error
	return promos, err
}

func (r *productPromoRepositoryImpl) FindById(id uint64) (entities.ProductPromo, error) {
	var promo entities.ProductPromo
	err := r.db.Order("id desc").Preload("Store").Preload("Product").First(&promo, id).Error
	return promo, err
}

func (r *productPromoRepositoryImpl) FindByProductId(productId uint32) ([]entities.ProductPromo, error) {
	var promos []entities.ProductPromo
	err := r.db.Order("id desc").Preload("Store").Preload("Product").Where("id_produk = ?", productId).Find(&promos).Error
	return promos, err
}

func (r *productPromoRepositoryImpl) Insert(promo models.ProductPromoRequest, storeId uint64) (entities.ProductPromo, error) {
	now := time.Now()
	newPromo := entities.ProductPromo{
		IDToko:    storeId,
		IDProduk:  promo.IDProduk,
		Promo:     promo.Promo,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	err := r.db.Create(&newPromo).Error
	if err != nil {
		return entities.ProductPromo{}, err
	}

	var completePromo entities.ProductPromo
	err = r.db.Preload("Store").Preload("Product").First(&completePromo, newPromo.ID).Error
	return completePromo, err
}

func (r *productPromoRepositoryImpl) Update(promo models.ProductPromoRequest, id uint64) (entities.ProductPromo, error) {
	var existingPromo entities.ProductPromo
	now := time.Now()

	err := r.db.First(&existingPromo, id).Error
	if err != nil {
		return entities.ProductPromo{}, err
	}

	existingPromo.Promo = promo.Promo
	existingPromo.IDToko = uint64(promo.IDToko)
	existingPromo.IDProduk = promo.IDProduk
	existingPromo.UpdatedAt = &now

	err = r.db.Save(&existingPromo).Error
	if err != nil {
		return entities.ProductPromo{}, err
	}

	var updatedPromo entities.ProductPromo
	err = r.db.Preload("Store").Preload("Product").First(&updatedPromo, id).Error
	return updatedPromo, err
}

func (r *productPromoRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entities.ProductPromo{}, id).Error
}

func (r *productPromoRepositoryImpl) DeleteAll() error {
	return r.db.Exec("DELETE FROM product_promos").Error
}
