package services

import (
	"errors"
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
)

// Contract
type AddressService interface {
	GetAll(user_id uint) ([]models.AddressResponse, error)
	GetById(id uint, user_id uint) (models.AddressResponse, error)
	Create(payload models.AddressRequest, user_id uint) (models.AddressResponse, error)
	Edit(id uint, payload models.AddressRequest, user_id uint) (models.AddressResponse, error)
	Delete(id uint, user_id uint) (models.AddressResponse, error)
}

type addressServiceImpl struct {
	repository     repositories.AddressRepository
	userRepository repositories.UserRepository
	regionService  RegionService
}

func NewAddressService(
	addressRepository *repositories.AddressRepository,
	userRepository *repositories.UserRepository,
	regionService *RegionService,
) AddressService {
	return &addressServiceImpl{
		repository:     *addressRepository,
		userRepository: *userRepository,
		regionService:  *regionService,
	}
}

func (service *addressServiceImpl) GetAll(user_id uint) ([]models.AddressResponse, error) {
	// Check if user exists and is admin
	user, err := service.userRepository.FindById(user_id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// Debug print
	fmt.Printf("User ID: %d, IsAdmin: %v\n", user.ID, user.IsAdmin)

	var addresses []entities.Address

	if user.IsAdmin {
		// Admin can see all addresses
		addresses, err = service.repository.FindAll()
	} else {
		// Regular users can only see their own addresses
		addresses, err = service.repository.FindByUserId(user_id)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to fetch addresses: %v", err)
	}

	// Debug print
	fmt.Printf("Found %d addresses\n", len(addresses))

	// mapping response
	responses := []models.AddressResponse{}

	for _, address := range addresses {
		// Get province and city data for each address
		provinceData, err := service.regionService.GetProvince(address.IDProvinsi)
		if err != nil {
			return nil, fmt.Errorf("failed to get province data: %v", err)
		}

		cityData, err := service.regionService.GetCity(address.IDKota)
		if err != nil {
			return nil, fmt.Errorf("failed to get city data: %v", err)
		}

		response := models.AddressResponse{
			ID:           address.ID,
			IDUser:       address.IDUser,
			JudulAlamat:  address.JudulAlamat,
			NamaPenerima: address.NamaPenerima,
			NoTelp:       address.NoTelp,
			DetailAlamat: address.DetailAlamat,
			Province: models.LocationResponse{
				ID:   provinceData.ID,
				Name: provinceData.Name,
			},
			City: models.CityResponse{
				ID:         cityData.ID,
				ProvinceID: cityData.ProvinceID,
				Name:       cityData.Name,
			},
			CreatedAt: address.CreatedAt,
			UpdatedAt: address.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *addressServiceImpl) GetById(id uint, user_id uint) (models.AddressResponse, error) {
	address, err := service.repository.FindById(id)
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Add debug prints
	fmt.Printf("Address user ID: %d, Requesting user ID: %d\n", address.IDUser, user_id)

	// Check if user is admin
	user, err := service.userRepository.FindById(user_id)
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Allow access if user is admin or owns the address
	if !user.IsAdmin && address.IDUser != user_id {
		return models.AddressResponse{}, errors.New("forbidden")
	}

	// Get province and city data
	province, err := service.regionService.GetProvince(address.IDProvinsi)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get province: %v", err)
	}

	city, err := service.regionService.GetCity(address.IDKota)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get city: %v", err)
	}

	response := models.AddressResponse{
		ID:           address.ID,
		IDUser:       address.IDUser,
		JudulAlamat:  address.JudulAlamat,
		NamaPenerima: address.NamaPenerima,
		NoTelp:       address.NoTelp,
		DetailAlamat: address.DetailAlamat,
		Province: models.LocationResponse{
			ID:   province.ID,
			Name: province.Name,
		},
		City: models.CityResponse{
			ID:         city.ID,
			ProvinceID: city.ProvinceID,
			Name:       city.Name,
		},
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}

	return response, nil
}

func (service *addressServiceImpl) Create(payload models.AddressRequest, user_id uint) (models.AddressResponse, error) {
	// If IDUser is not provided in the request, use the authenticated user's ID
	targetUserId := payload.IDUser
	if targetUserId == 0 {
		targetUserId = user_id
	}

	// Verify user exists
	_, err := service.userRepository.FindById(targetUserId)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("user not found: %v", err)
	}

	// Create address
	address := entities.Address{
		IDUser:       targetUserId,
		JudulAlamat:  payload.JudulAlamat,
		NamaPenerima: payload.NamaPenerima,
		NoTelp:       payload.NoTelp,
		DetailAlamat: payload.DetailAlamat,
		IDProvinsi:   payload.IDProvinsi,
		IDKota:       payload.IDKota,
	}

	success, err := service.repository.Insert(address)
	if err != nil || !success {
		return models.AddressResponse{}, fmt.Errorf("failed to create address: %v", err)
	}

	// Fetch the created address
	created_address, err := service.repository.FindByCondition(map[string]interface{}{
		"id_user":       targetUserId,
		"judul_alamat":  payload.JudulAlamat,
		"nama_penerima": payload.NamaPenerima,
		"no_telp":       payload.NoTelp,
		"detail_alamat": payload.DetailAlamat,
	})
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Get province and city data from region service
	provinceData, err := service.regionService.GetProvince(payload.IDProvinsi)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get province: %v", err)
	}

	cityData, err := service.regionService.GetCity(payload.IDKota)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get city: %v", err)
	}

	// Map the region data to our response types
	province := models.LocationResponse{
		ID:   provinceData.ID,
		Name: provinceData.Name,
	}

	city := models.CityResponse{
		ID:         cityData.ID,
		ProvinceID: cityData.ProvinceID,
		Name:       cityData.Name,
	}

	response := models.AddressResponse{
		ID:           created_address.ID,
		IDUser:       created_address.IDUser,
		JudulAlamat:  created_address.JudulAlamat,
		NamaPenerima: created_address.NamaPenerima,
		NoTelp:       created_address.NoTelp,
		DetailAlamat: created_address.DetailAlamat,
		Province:     province,
		City:         city,
		CreatedAt:    created_address.CreatedAt,
		UpdatedAt:    created_address.UpdatedAt,
	}

	return response, nil
}

func (service *addressServiceImpl) Edit(id uint, payload models.AddressRequest, user_id uint) (models.AddressResponse, error) {
	// Check if address exists
	check_address, err := service.repository.FindById(id)
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Check if user is admin
	user, err := service.userRepository.FindById(user_id)
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Allow access if user is admin or owns the address
	if !user.IsAdmin && check_address.IDUser != user_id {
		return models.AddressResponse{}, errors.New("forbidden")
	}

	// Create address entity
	address := entities.Address{
		IDUser:       check_address.IDUser, // Keep original user ID
		JudulAlamat:  check_address.JudulAlamat,
		NamaPenerima: payload.NamaPenerima,
		NoTelp:       payload.NoTelp,
		DetailAlamat: payload.DetailAlamat,
		IDProvinsi:   payload.IDProvinsi,
		IDKota:       payload.IDKota,
	}

	// Update address
	success, err := service.repository.Update(id, address)
	if err != nil || !success {
		return models.AddressResponse{}, fmt.Errorf("failed to update address: %v", err)
	}

	// Fetch updated address
	updated_address, err := service.repository.FindById(id)
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Get province and city data
	provinceData, err := service.regionService.GetProvince(updated_address.IDProvinsi)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get province: %v", err)
	}

	cityData, err := service.regionService.GetCity(updated_address.IDKota)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get city: %v", err)
	}

	// Create response
	response := models.AddressResponse{
		ID:           updated_address.ID,
		IDUser:       updated_address.IDUser,
		JudulAlamat:  updated_address.JudulAlamat,
		NamaPenerima: updated_address.NamaPenerima,
		NoTelp:       updated_address.NoTelp,
		DetailAlamat: updated_address.DetailAlamat,
		Province: models.LocationResponse{
			ID:   provinceData.ID,
			Name: provinceData.Name,
		},
		City: models.CityResponse{
			ID:         cityData.ID,
			ProvinceID: cityData.ProvinceID,
			Name:       cityData.Name,
		},
		CreatedAt: updated_address.CreatedAt,
		UpdatedAt: updated_address.UpdatedAt,
	}

	return response, nil
}

func (service *addressServiceImpl) Delete(id uint, user_id uint) (models.AddressResponse, error) {
	// Get address data before deletion
	address, err := service.repository.FindById(id)
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Check if user is admin
	user, err := service.userRepository.FindById(user_id)
	if err != nil {
		return models.AddressResponse{}, err
	}

	// Allow access if user is admin or owns the address
	if !user.IsAdmin && address.IDUser != user_id {
		return models.AddressResponse{}, errors.New("forbidden")
	}

	// Get province and city data before deletion
	provinceData, err := service.regionService.GetProvince(address.IDProvinsi)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get province data: %v", err)
	}

	cityData, err := service.regionService.GetCity(address.IDKota)
	if err != nil {
		return models.AddressResponse{}, fmt.Errorf("failed to get city data: %v", err)
	}

	// Delete the address
	success, err := service.repository.Destroy(id)
	if err != nil || !success {
		return models.AddressResponse{}, err
	}

	// Create response with complete data
	response := models.AddressResponse{
		ID:           address.ID,
		IDUser:       address.IDUser,
		JudulAlamat:  address.JudulAlamat,
		NamaPenerima: address.NamaPenerima,
		NoTelp:       address.NoTelp,
		DetailAlamat: address.DetailAlamat,
		Province: models.LocationResponse{
			ID:   provinceData.ID,
			Name: provinceData.Name,
		},
		City: models.CityResponse{
			ID:         cityData.ID,
			ProvinceID: cityData.ProvinceID,
			Name:       cityData.Name,
		},
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}

	return response, nil
}
