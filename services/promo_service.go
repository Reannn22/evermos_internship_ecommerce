package services

import (
	"errors"
	"mini-project-evermos/models"
	"mini-project-evermos/repositories"
)

type ProductPromoService interface {
	GetAll() ([]models.ProductPromoResponse, error)
	GetById(id uint64) (models.ProductPromoResponse, error)
	GetByProductId(productId uint32) ([]models.ProductPromoResponse, error)
	Create(input models.ProductPromoRequest) (models.ProductPromoResponse, error)
	Update(input models.ProductPromoRequest, id uint64, userId uint64) (models.ProductPromoResponse, error)
	Delete(id uint64, userId uint64) error
	ClearAll() ([]models.ProductPromoResponse, error)
}

type productPromoServiceImpl struct {
	promoRepository repositories.ProductPromoRepository
}

func NewProductPromoService(promoRepository repositories.ProductPromoRepository) ProductPromoService {
	return &productPromoServiceImpl{promoRepository: promoRepository}
}

func (s *productPromoServiceImpl) GetAll() ([]models.ProductPromoResponse, error) {
	promos, err := s.promoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.ProductPromoResponse
	for _, promo := range promos {
		responses = append(responses, models.ProductPromoResponse{
			ID:        uint(promo.ID),
			IDToko:    uint(promo.IDToko),
			IDProduk:  uint(promo.IDProduk),
			Promo:     promo.Promo,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *productPromoServiceImpl) GetById(id uint64) (models.ProductPromoResponse, error) {
	promo, err := s.promoRepository.FindById(id)
	if err != nil {
		return models.ProductPromoResponse{}, err
	}

	return models.ProductPromoResponse{
		ID:        uint(promo.ID),
		IDToko:    uint(promo.IDToko),
		IDProduk:  uint(promo.IDProduk),
		Promo:     promo.Promo,
		CreatedAt: promo.CreatedAt,
		UpdatedAt: promo.UpdatedAt,
	}, nil
}

func (s *productPromoServiceImpl) GetByProductId(productId uint32) ([]models.ProductPromoResponse, error) {
	promos, err := s.promoRepository.FindByProductId(productId)
	if err != nil {
		return nil, err
	}

	var responses []models.ProductPromoResponse
	for _, promo := range promos {
		responses = append(responses, models.ProductPromoResponse{
			ID:        uint(promo.ID),
			IDToko:    uint(promo.IDToko),
			IDProduk:  uint(promo.IDProduk),
			Promo:     promo.Promo,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *productPromoServiceImpl) Create(input models.ProductPromoRequest) (models.ProductPromoResponse, error) {
	promo, err := s.promoRepository.Insert(input, uint64(input.IDToko))
	if err != nil {
		return models.ProductPromoResponse{}, err
	}

	return s.GetById(promo.ID)
}

func (s *productPromoServiceImpl) Update(input models.ProductPromoRequest, id uint64, userId uint64) (models.ProductPromoResponse, error) {
	existingPromo, err := s.promoRepository.FindById(id)
	if err != nil {
		return models.ProductPromoResponse{}, err
	}

	if existingPromo.ID == 0 {
		return models.ProductPromoResponse{}, errors.New("promo not found")
	}

	updatedPromo, err := s.promoRepository.Update(input, id)
	if err != nil {
		return models.ProductPromoResponse{}, err
	}

	return s.GetById(updatedPromo.ID)
}

func (s *productPromoServiceImpl) Delete(id uint64, userId uint64) error {
	return s.promoRepository.Delete(id)
}

func (s *productPromoServiceImpl) ClearAll() ([]models.ProductPromoResponse, error) {
	// First get all promos
	promos, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	// Then delete all promos
	err = s.promoRepository.DeleteAll()
	if err != nil {
		return nil, err
	}

	return promos, nil
}
