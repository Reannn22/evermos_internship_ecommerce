package repositories

import (
	"errors"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

type KeranjangBelanjaRepository interface {
	FindAll() ([]entities.KeranjangBelanja, error)
	FindById(id uint) (entities.KeranjangBelanja, error)
	Create(input models.KeranjangBelanjaRequest) (entities.KeranjangBelanja, error)
	Update(id uint, input models.KeranjangBelanjaRequest) (entities.KeranjangBelanja, error)
	Delete(id uint) (entities.KeranjangBelanja, error)
	ClearAll() error
}

type keranjangBelanjaRepositoryImpl struct {
	db *gorm.DB
}

func NewKeranjangBelanjaRepository(db *gorm.DB) KeranjangBelanjaRepository {
	return &keranjangBelanjaRepositoryImpl{db}
}

func (repository *keranjangBelanjaRepositoryImpl) FindAll() ([]entities.KeranjangBelanja, error) {
	var keranjangBelanja []entities.KeranjangBelanja
	err := repository.db.Order("id desc").Preload("Store.FotoToko").Preload("Product").Preload("Product.FotoProduk").Find(&keranjangBelanja).Error
	if err != nil {
		return nil, err
	}

	// Update UrlFoto for each Store if FotoToko exists
	for i := range keranjangBelanja {
		if len(keranjangBelanja[i].Store.FotoToko) > 0 {
			keranjangBelanja[i].Store.UrlFoto = keranjangBelanja[i].Store.FotoToko[0].URL
		}
	}

	return keranjangBelanja, nil
}

func (repository *keranjangBelanjaRepositoryImpl) FindById(id uint) (entities.KeranjangBelanja, error) {
	var keranjangBelanja entities.KeranjangBelanja
	err := repository.db.
		Order("id desc").
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at")
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.FotoProduk").
		First(&keranjangBelanja, id).Error
	if err != nil {
		return keranjangBelanja, err
	}

	// Update UrlFoto if FotoToko exists
	if len(keranjangBelanja.Store.FotoToko) > 0 {
		keranjangBelanja.Store.UrlFoto = keranjangBelanja.Store.FotoToko[0].URL
	}

	return keranjangBelanja, nil
}

func (repository *keranjangBelanjaRepositoryImpl) Create(input models.KeranjangBelanjaRequest) (entities.KeranjangBelanja, error) {
	// Verify store exists
	var store entities.Store
	if err := repository.db.First(&store, input.IDToko).Error; err != nil {
		return entities.KeranjangBelanja{}, errors.New("store not found")
	}

	// Verify product exists
	var product entities.Product
	if err := repository.db.First(&product, input.IDProduk).Error; err != nil {
		return entities.KeranjangBelanja{}, errors.New("product not found")
	}

	keranjangBelanja := entities.KeranjangBelanja{
		IDToko:   input.IDToko,
		IDProduk: input.IDProduk,
	}

	err := repository.db.Create(&keranjangBelanja).Error
	if err != nil {
		return entities.KeranjangBelanja{}, err
	}

	// Reload the data with associations
	var completeKeranjang entities.KeranjangBelanja
	err = repository.db.
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.FotoProduk").
		First(&completeKeranjang, keranjangBelanja.ID).Error
	if err != nil {
		return completeKeranjang, err
	}

	// Update UrlFoto if FotoToko exists
	if len(completeKeranjang.Store.FotoToko) > 0 {
		completeKeranjang.Store.UrlFoto = completeKeranjang.Store.FotoToko[0].URL
	}

	return completeKeranjang, nil
}

func (repository *keranjangBelanjaRepositoryImpl) Update(id uint, input models.KeranjangBelanjaRequest) (entities.KeranjangBelanja, error) {
	// Verify the cart item exists
	var keranjangBelanja entities.KeranjangBelanja
	if err := repository.db.First(&keranjangBelanja, id).Error; err != nil {
		return entities.KeranjangBelanja{}, errors.New("cart item not found")
	}

	// Verify store exists
	var store entities.Store
	if err := repository.db.First(&store, input.IDToko).Error; err != nil {
		return entities.KeranjangBelanja{}, errors.New("store not found")
	}

	// Verify product exists
	var product entities.Product
	if err := repository.db.First(&product, input.IDProduk).Error; err != nil {
		return entities.KeranjangBelanja{}, errors.New("product not found")
	}

	// Update the cart item
	keranjangBelanja.IDToko = input.IDToko
	keranjangBelanja.IDProduk = input.IDProduk
	keranjangBelanja.JumlahProduk = input.JumlahProduk

	err := repository.db.Save(&keranjangBelanja).Error
	if err != nil {
		return entities.KeranjangBelanja{}, err
	}

	// Reload the data with associations
	err = repository.db.Preload("Store").Preload("Product").First(&keranjangBelanja, keranjangBelanja.ID).Error
	return keranjangBelanja, err
}

func (repository *keranjangBelanjaRepositoryImpl) Delete(id uint) (entities.KeranjangBelanja, error) {
	var keranjangBelanja entities.KeranjangBelanja

	// Get complete data with associations before deletion
	err := repository.db.Preload("Store").Preload("Product").First(&keranjangBelanja, id).Error
	if err != nil {
		return keranjangBelanja, err
	}

	// Delete the record
	err = repository.db.Delete(&keranjangBelanja).Error
	return keranjangBelanja, err
}

func (repository *keranjangBelanjaRepositoryImpl) ClearAll() error {
	return repository.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entities.KeranjangBelanja{}).Error
}
