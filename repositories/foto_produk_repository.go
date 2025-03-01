package repositories

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

type FotoProdukRepository interface {
	Create(request models.FotoProdukRequest) (entities.FotoProduk, error)
	FindById(id uint) (entities.FotoProduk, error)
	Update(id uint, request models.FotoProdukRequest) (entities.FotoProduk, error)
	Delete(id uint) error
	FindAll() ([]entities.FotoProduk, error)
	FindByProductId(productId uint) ([]entities.FotoProduk, error)
	CheckProductExists(productID uint) (bool, error)
}

type fotoProdukRepositoryImpl struct {
	database *gorm.DB
}

func NewFotoProdukRepository(database *gorm.DB) FotoProdukRepository {
	return &fotoProdukRepositoryImpl{database}
}

func (repository *fotoProdukRepositoryImpl) Create(request models.FotoProdukRequest) (entities.FotoProduk, error) {
	foto := entities.FotoProduk{
		IDProduk: request.ProductID,
		PhotoURL: request.PhotoURL,
		URL:      request.PhotoURL,
		Photo:    "/uploads/products/" + request.File.Filename,
	}

	if err := repository.database.Create(&foto).Error; err != nil {
		return entities.FotoProduk{}, err
	}

	return foto, nil
}

func (repository *fotoProdukRepositoryImpl) FindById(id uint) (entities.FotoProduk, error) {
	var foto entities.FotoProduk
	err := repository.database.First(&foto, id).Error
	return foto, err
}

func (repository *fotoProdukRepositoryImpl) Update(id uint, request models.FotoProdukRequest) (entities.FotoProduk, error) {
	foto := entities.FotoProduk{
		ID:       id,
		IDProduk: request.ProductID,
		PhotoURL: request.PhotoURL,
		URL:      request.PhotoURL,
	}

	if request.File != nil {
		foto.Photo = "/uploads/products/" + request.File.Filename
	}

	if err := repository.database.Model(&foto).Updates(foto).Error; err != nil {
		return entities.FotoProduk{}, err
	}

	return foto, nil
}

func (repository *fotoProdukRepositoryImpl) Delete(id uint) error {
	return repository.database.Delete(&entities.FotoProduk{}, id).Error
}

func (repository *fotoProdukRepositoryImpl) FindAll() ([]entities.FotoProduk, error) {
	var photos []entities.FotoProduk
	err := repository.database.Order("id desc").Find(&photos).Error
	return photos, err
}

func (repository *fotoProdukRepositoryImpl) FindByProductId(productId uint) ([]entities.FotoProduk, error) {
	var photos []entities.FotoProduk
	err := repository.database.Where("id_produk = ?", productId).Order("id desc").Find(&photos).Error
	return photos, err
}

func (repository *fotoProdukRepositoryImpl) CheckProductExists(productID uint) (bool, error) {
	var count int64
	err := repository.database.Model(&entities.Product{}).Where("id = ?", productID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
