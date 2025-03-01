package region

import (
	"encoding/json"
	"io/ioutil"
	"mini-project-evermos/models"
	"net/http"
)

const BASE_URL = "https://emsifa.github.io/api-wilayah-indonesia/api"

func GetProvinceByID(province_id string) (*models.Province, error) {
	response, err := http.Get(BASE_URL + "/province/" + province_id + ".json")

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject models.Province
	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}

func GetCityByID(city_id string) (*models.City, error) {
	response, err := http.Get(BASE_URL + "/regency/" + city_id + ".json")

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject models.City
	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}

func GetAllProvince() ([]models.Province, error) {
	response, err := http.Get(BASE_URL + "/provinces.json")

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []models.Province
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

func GetAllCity(province_id string) ([]models.City, error) {
	response, err := http.Get(BASE_URL + "/regencies/" + province_id + ".json")

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject []models.City
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}
