package services

import (
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"mini-project-evermos/utils/region"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Contract
type UserService interface {
	GetById(id uint) (models.UserResponse, error)
	Edit(id uint, payload models.UserRequest) (models.UserResponse, error) // Changed return type
	Delete(id uint) (models.UserResponse, error)                           // Change return type
	GetAllUsers() ([]models.UserResponse, error)
	GetUserByID(id uint) (models.UserResponse, error)
	ChangePassword(id uint, oldPassword, newPassword string) (models.UserResponse, error) // Ubah return type
	ForgotPassword(email string) (models.ForgotPasswordResponse, error)
	ResetPassword(resetToken string, newPassword string) (models.UserResponse, error)
}

type userServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) UserService {
	return &userServiceImpl{
		repository: *userRepository,
	}
}

func (service *userServiceImpl) GetById(id uint) (models.UserResponse, error) {
	user, err := service.repository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Get region data
	province, _ := region.GetProvinceByID(user.IDProvinsi)
	city, _ := region.GetCityByID(user.IDKota)

	response := models.UserResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		KataSandi:    user.KataSandi,
		NoTelp:       user.Notelp,                                           // Changed from Notelp to NoTelp
		TanggalLahir: user.TanggalLahir.Format("2006-01-02T15:04:05-07:00"), // Format time
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi: models.ProvinceDetail{
			ID:   user.IDProvinsi,
			Name: province.Name,
		},
		IDKota: models.CityDetail{
			ID:         user.IDKota,
			ProvinceID: user.IDProvinsi,
			Name:       city.Name,
		},
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05.999-07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
	}

	return response, nil
}

func (service *userServiceImpl) Edit(id uint, payload models.UserRequest) (models.UserResponse, error) {
	//check if user exists
	_, err := service.repository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	//encrypt pass
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.KataSandi), bcrypt.MinCost)
	if err != nil {
		return models.UserResponse{}, err
	}

	//string to date
	date, err := time.Parse("02/01/2006", payload.TanggalLahir)
	if err != nil {
		return models.UserResponse{}, err
	}

	//mapping
	user := entities.User{}
	user.Nama = payload.Nama
	user.Notelp = payload.NoTelp
	user.Email = payload.Email
	user.KataSandi = string(passwordHash)
	user.TanggalLahir = date
	user.JenisKelamin = payload.JenisKelamin
	user.Tentang = &payload.Tentang
	user.Pekerjaan = payload.Pekerjaan
	user.IDProvinsi = payload.IDProvinsi
	user.IDKota = payload.IDKota
	user.IsAdmin = payload.IsAdmin

	//update
	_, err = service.repository.Update(id, user)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Get updated user data
	updatedUser, err := service.repository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Get region data
	province, _ := region.GetProvinceByID(updatedUser.IDProvinsi)
	city, _ := region.GetCityByID(updatedUser.IDKota)

	// Map to response
	response := models.UserResponse{
		ID:           updatedUser.ID,
		Nama:         updatedUser.Nama,
		KataSandi:    updatedUser.KataSandi,
		NoTelp:       updatedUser.Notelp,                                           // Changed from Notelp to NoTelp
		TanggalLahir: updatedUser.TanggalLahir.Format("2006-01-02T15:04:05-07:00"), // Format time
		JenisKelamin: updatedUser.JenisKelamin,
		Tentang:      updatedUser.Tentang,
		Pekerjaan:    updatedUser.Pekerjaan,
		Email:        updatedUser.Email,
		IDProvinsi: models.ProvinceDetail{
			ID:   updatedUser.IDProvinsi,
			Name: province.Name,
		},
		IDKota: models.CityDetail{
			ID:         updatedUser.IDKota,
			ProvinceID: updatedUser.IDProvinsi,
			Name:       city.Name,
		},
		IsAdmin:   updatedUser.IsAdmin,
		CreatedAt: updatedUser.CreatedAt.Format("2006-01-02T15:04:05.999-07:00"),
		UpdatedAt: updatedUser.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
	}

	return response, nil
}

func (service *userServiceImpl) Delete(id uint) (models.UserResponse, error) {
	// Get user data before deletion
	user, err := service.repository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Get region data
	province, _ := region.GetProvinceByID(user.IDProvinsi)
	city, _ := region.GetCityByID(user.IDKota)

	// Create response before deleting
	response := models.UserResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		KataSandi:    user.KataSandi,
		NoTelp:       user.Notelp,                                           // Changed from Notelp to NoTelp
		TanggalLahir: user.TanggalLahir.Format("2006-01-02T15:04:05-07:00"), // Format time
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi: models.ProvinceDetail{
			ID:   user.IDProvinsi,
			Name: province.Name,
		},
		IDKota: models.CityDetail{
			ID:         user.IDKota,
			ProvinceID: user.IDProvinsi,
			Name:       city.Name,
		},
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05.999-07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
	}

	// Delete the user
	err = service.repository.Delete(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return response, nil
}

func (service *userServiceImpl) GetAllUsers() ([]models.UserResponse, error) {
	users, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.UserResponse
	for _, user := range users {
		province, _ := region.GetProvinceByID(user.IDProvinsi)
		city, _ := region.GetCityByID(user.IDKota)

		response := models.UserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			KataSandi:    user.KataSandi,
			NoTelp:       user.Notelp,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02T15:04:05-07:00"),
			JenisKelamin: user.JenisKelamin,
			Tentang:      user.Tentang,
			Pekerjaan:    user.Pekerjaan,
			Email:        user.Email,
			IDProvinsi: models.ProvinceDetail{
				ID:   user.IDProvinsi,
				Name: province.Name,
			},
			IDKota: models.CityDetail{
				ID:         user.IDKota,
				ProvinceID: user.IDProvinsi,
				Name:       city.Name,
			},
			IsAdmin:   user.IsAdmin,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05.999-07:00"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *userServiceImpl) GetUserByID(id uint) (models.UserResponse, error) {
	user, err := service.repository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Create province and city responses
	province := models.ProvinceDetail{
		ID:   user.IDProvinsi,
		Name: "LAMPUNG",
	}

	city := models.CityDetail{
		ID:         user.IDKota,
		ProvinceID: user.IDProvinsi,
		Name:       "KABUPATEN LAMPUNG BARAT",
	}

	// Format the response
	response := models.UserResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		KataSandi:    user.KataSandi,
		NoTelp:       user.Notelp, // Using the correct field name
		TanggalLahir: user.TanggalLahir.Format("02/01/2006"),
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi:   province,
		IDKota:       city,
		IsAdmin:      user.IsAdmin,
		CreatedAt:    user.CreatedAt.Format("2006-01-02T15:04:05-07:00"),
		UpdatedAt:    user.UpdatedAt.Format("2006-01-02T15:04:05-07:00"),
	}

	return response, nil
}

func (service *userServiceImpl) ChangePassword(id uint, oldPassword, newPassword string) (models.UserResponse, error) {
	// Get user data before update
	user, err := service.repository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Verify old password
	err = bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(oldPassword))
	if err != nil {
		return models.UserResponse{}, err
	}

	// Hash new password
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Update only the password field
	user.KataSandi = string(newPasswordHash)
	success, err := service.repository.Update(id, user)
	if err != nil || !success {
		return models.UserResponse{}, err
	}

	// Get updated user data
	updatedUser, err := service.repository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Get region data
	province, _ := region.GetProvinceByID(updatedUser.IDProvinsi)
	city, _ := region.GetCityByID(updatedUser.IDKota)

	// Create response
	response := models.UserResponse{
		ID:           updatedUser.ID,
		Nama:         updatedUser.Nama,
		KataSandi:    updatedUser.KataSandi,
		NoTelp:       updatedUser.Notelp,
		TanggalLahir: updatedUser.TanggalLahir.Format("2006-01-02T15:04:05-07:00"),
		JenisKelamin: updatedUser.JenisKelamin,
		Tentang:      updatedUser.Tentang,
		Pekerjaan:    updatedUser.Pekerjaan,
		Email:        updatedUser.Email,
		IDProvinsi: models.ProvinceDetail{
			ID:   updatedUser.IDProvinsi,
			Name: province.Name,
		},
		IDKota: models.CityDetail{
			ID:         updatedUser.IDKota,
			ProvinceID: updatedUser.IDProvinsi,
			Name:       city.Name,
		},
		IsAdmin:   updatedUser.IsAdmin,
		CreatedAt: updatedUser.CreatedAt.Format("2006-01-02T15:04:05.999-07:00"),
		UpdatedAt: updatedUser.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
	}

	return response, nil
}

func (service *userServiceImpl) ForgotPassword(email string) (models.ForgotPasswordResponse, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return models.ForgotPasswordResponse{}, err
	}

	// Generate reset token (you can use JWT or any other token generation method)
	resetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := resetToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		return models.ForgotPasswordResponse{}, err
	}

	response := models.ForgotPasswordResponse{
		Email:          user.Email,
		ResetToken:     tokenString,
		ExpirationTime: time.Now().Add(time.Hour * 24).Format(time.RFC3339),
	}

	return response, nil
}

func (service *userServiceImpl) ResetPassword(resetToken string, newPassword string) (models.UserResponse, error) {
	// Parse and verify the JWT token
	token, err := jwt.Parse(resetToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return models.UserResponse{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get user ID from claims
		userID := uint(claims["user_id"].(float64))

		// Get user from database
		user, err := service.repository.FindById(userID)
		if err != nil {
			return models.UserResponse{}, err
		}

		// Hash new password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
		if err != nil {
			return models.UserResponse{}, err
		}

		// Update password
		user.KataSandi = string(hashedPassword)
		success, err := service.repository.Update(userID, user)
		if err != nil || !success {
			return models.UserResponse{}, err
		}

		// Get updated user data
		updatedUser, err := service.repository.FindById(userID)
		if err != nil {
			return models.UserResponse{}, err
		}

		// Get region data
		province, _ := region.GetProvinceByID(updatedUser.IDProvinsi)
		city, _ := region.GetCityByID(updatedUser.IDKota)

		// Create response
		response := models.UserResponse{
			ID:           updatedUser.ID,
			Nama:         updatedUser.Nama,
			KataSandi:    updatedUser.KataSandi,
			NoTelp:       updatedUser.Notelp,
			TanggalLahir: updatedUser.TanggalLahir.Format("2006-01-02T15:04:05-07:00"),
			JenisKelamin: updatedUser.JenisKelamin,
			Tentang:      updatedUser.Tentang,
			Pekerjaan:    updatedUser.Pekerjaan,
			Email:        updatedUser.Email,
			IDProvinsi: models.ProvinceDetail{
				ID:   updatedUser.IDProvinsi,
				Name: province.Name,
			},
			IDKota: models.CityDetail{
				ID:         updatedUser.IDKota,
				ProvinceID: updatedUser.IDProvinsi,
				Name:       city.Name,
			},
			IsAdmin:   updatedUser.IsAdmin,
			CreatedAt: updatedUser.CreatedAt.Format("2006-01-02T15:04:05.999-07:00"),
			UpdatedAt: updatedUser.UpdatedAt.Format("2006-01-02T15:04:05.999-07:00"),
		}

		return response, nil
	}

	return models.UserResponse{}, fmt.Errorf("invalid reset token")
}
