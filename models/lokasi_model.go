package models

// Shared location types used across different models
type ProvinceDetail struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CityDetail struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// LocationResponse and CityResponse are now the shared types for all responses
type LocationResponse = ProvinceDetail // Using type alias
type CityResponse = CityDetail         // Using type alias
