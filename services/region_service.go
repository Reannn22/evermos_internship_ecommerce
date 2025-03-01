package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/utils/region"
)

// Contract
type RegionService interface {
	GetAllProvince() ([]models.Province, error)
	GetProvince(prov_id string) (*models.Province, error)
	GetAllCity(prov_id string) ([]models.City, error)
	GetCity(city_id string) (*models.City, error)
}

type Province struct {
	ID   string
	Name string
}

type City struct {
	ID         string
	ProvinceID string
	Name       string
}

type regionServiceImpl struct {
}

func NewRegionService() RegionService {
	return &regionServiceImpl{}
}

func (service *regionServiceImpl) GetAllProvince() ([]models.Province, error) {
	province, err := region.GetAllProvince()

	return province, err
}

func (service *regionServiceImpl) GetProvince(prov_id string) (*models.Province, error) {
	province, err := region.GetProvinceByID(prov_id)

	return province, err
}

func (service *regionServiceImpl) GetAllCity(prov_id string) ([]models.City, error) {
	city, err := region.GetAllCity(prov_id)

	return city, err
}

func (service *regionServiceImpl) GetCity(city_id string) (*models.City, error) {
	city, err := region.GetCityByID(city_id)
	return city, err
}
