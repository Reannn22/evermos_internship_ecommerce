package services

import (
	"errors"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"mini-project-evermos/utils/jwt"
	"mini-project-evermos/utils/region"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Contract
type AuthService interface {
	Register(input models.RegisterRequest) (models.RegisterResponse, error)
	Login(input models.LoginRequest) (models.LoginResponse, error)
	Logout() (models.LogoutResponse, error) // Add this
}

type authServiceImpl struct {
	repository     repositories.AuthRepository
	repositoryUser repositories.UserRepository
}

func NewAuthService(authRepository *repositories.AuthRepository, userRepository *repositories.UserRepository) AuthService {
	return &authServiceImpl{
		repository:     *authRepository,
		repositoryUser: *userRepository,
	}
}

func (service *authServiceImpl) Register(input models.RegisterRequest) (models.RegisterResponse, error) {

	//encrypt pass
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.KataSandi), bcrypt.MinCost)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	//string to date
	date, err := time.Parse("02/01/2006", input.TanggalLahir)

	if err != nil {
		return models.RegisterResponse{}, err
	}

	//mapping
	user := entities.User{}
	user.Nama = input.Nama
	user.Username = input.Email // Set username to email as default
	user.Notelp = input.NoTelp
	user.Email = input.Email
	user.KataSandi = string(passwordHash)
	user.TanggalLahir = date
	user.JenisKelamin = input.JenisKelamin // Added this
	user.Tentang = &input.Tentang          // Added this with pointer
	user.Pekerjaan = input.Pekerjaan
	user.IDProvinsi = input.IDProvinsi
	user.IDKota = input.IDKota
	user.IsAdmin = input.IsAdmin // Added this

	//register user
	newUser, err := service.repository.Register(user)

	if err != nil {
		return models.RegisterResponse{}, err
	}

	//get region data
	province, _ := region.GetProvinceByID(newUser.IDProvinsi)
	city, _ := region.GetCityByID(newUser.IDKota)

	// Map to response with formatted data
	response := models.RegisterResponse{
		ID:           newUser.ID,
		Nama:         newUser.Nama,
		KataSandi:    newUser.KataSandi,
		NoTelp:       newUser.Notelp,
		TanggalLahir: newUser.TanggalLahir.Format("02/01/2006"),
		JenisKelamin: newUser.JenisKelamin,
		Tentang:      newUser.Tentang,
		Pekerjaan:    newUser.Pekerjaan,
		Email:        newUser.Email,
		IDProvinsi: models.ProvinceDetail{
			ID:   newUser.IDProvinsi,
			Name: province.Name,
		},
		IDKota: models.CityDetail{
			ID:         newUser.IDKota,
			ProvinceID: newUser.IDProvinsi,
			Name:       city.Name,
		},
		IsAdmin:   newUser.IsAdmin,
		CreatedAt: newUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: newUser.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}

func (service *authServiceImpl) Login(input models.LoginRequest) (models.LoginResponse, error) {
	email := input.Email
	password := input.KataSandi

	//check user
	check_user, err := service.repositoryUser.FindByEmail(email)

	if err != nil {
		return models.LoginResponse{}, errors.New("Email Not Found")
	}

	//check login
	err = bcrypt.CompareHashAndPassword([]byte(check_user.KataSandi), []byte(password))

	if err != nil {
		return models.LoginResponse{}, errors.New("Email atau kata sandi salah")
	}

	//generate token jwt
	token, err := jwt.GenerateNewAccessToken(check_user)

	//get region
	province, err := region.GetProvinceByID(check_user.IDProvinsi)
	city, err := region.GetCityByID(check_user.IDKota)

	//response mapping
	response := models.LoginResponse{
		ID:           check_user.ID,
		Nama:         check_user.Nama,
		KataSandi:    check_user.KataSandi,
		NoTelp:       check_user.Notelp,
		TanggalLahir: check_user.TanggalLahir.Format("02/01/2006"),
		JenisKelamin: check_user.JenisKelamin,
		Tentang:      check_user.Tentang,
		Pekerjaan:    check_user.Pekerjaan,
		Email:        check_user.Email,
		IDProvinsi: models.ProvinceDetail{
			ID:   check_user.IDProvinsi,
			Name: province.Name,
		},
		IDKota: models.CityDetail{
			ID:         check_user.IDKota,
			ProvinceID: check_user.IDProvinsi,
			Name:       city.Name,
		},
		IsAdmin:   check_user.IsAdmin,
		CreatedAt: check_user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: check_user.UpdatedAt.Format(time.RFC3339),
		Token:     token,
	}

	return response, nil
}

func (service *authServiceImpl) Logout() (models.LogoutResponse, error) {
	// Get the last logged in user (you may want to modify this to get the actual logged-in user)
	user, err := service.repositoryUser.FindLastUser()
	if err != nil {
		return models.LogoutResponse{}, err
	}

	// Get region data
	province, _ := region.GetProvinceByID(user.IDProvinsi)
	city, _ := region.GetCityByID(user.IDKota)

	// Map to response
	response := models.LogoutResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		KataSandi:    user.KataSandi,
		NoTelp:       user.Notelp,
		TanggalLahir: user.TanggalLahir.Format("02/01/2006"),
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
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}
