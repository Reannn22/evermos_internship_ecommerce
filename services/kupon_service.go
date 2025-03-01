package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/repositories"
)

type ProductCouponService interface {
	GetAll() ([]models.ProductCouponResponse, error)
	GetById(id uint) (models.ProductCouponResponse, error)
	GetByCode(code string) (models.ProductCouponResponse, error)
	GetByProduct(productId uint) ([]models.ProductCouponResponse, error)
	Create(request models.ProductCouponRequest) (models.ProductCouponResponse, error)
	Update(id uint, request models.ProductCouponRequest) (models.ProductCouponResponse, error)
	Delete(id uint) (models.ProductCouponResponse, error)
	ValidateCoupon(code string, productId uint) (models.ProductCouponResponse, error)
}

type productCouponServiceImpl struct {
	repository repositories.ProductCouponRepository
}

func NewProductCouponService(repository repositories.ProductCouponRepository) ProductCouponService {
	return &productCouponServiceImpl{repository}
}

func (s *productCouponServiceImpl) GetAll() ([]models.ProductCouponResponse, error) {
	return s.repository.FindAll()
}

func (s *productCouponServiceImpl) GetById(id uint) (models.ProductCouponResponse, error) {
	return s.repository.FindById(id)
}

func (s *productCouponServiceImpl) GetByCode(code string) (models.ProductCouponResponse, error) {
	return s.repository.FindByCode(code)
}

func (s *productCouponServiceImpl) GetByProduct(productId uint) ([]models.ProductCouponResponse, error) {
	return s.repository.FindByProduct(productId)
}

func (s *productCouponServiceImpl) Create(request models.ProductCouponRequest) (models.ProductCouponResponse, error) {
	return s.repository.Create(request)
}

func (s *productCouponServiceImpl) Update(id uint, request models.ProductCouponRequest) (models.ProductCouponResponse, error) {
	return s.repository.Update(id, request)
}

func (s *productCouponServiceImpl) Delete(id uint) (models.ProductCouponResponse, error) {
	return s.repository.Delete(id)
}

func (s *productCouponServiceImpl) ValidateCoupon(code string, productId uint) (models.ProductCouponResponse, error) {
	return s.repository.ValidateCoupon(code, productId)
}
