package formatter

import "mini-project-evermos/models"

func FormatStore(store models.StoreResponse) models.StoreResponse {
	return models.StoreResponse{
		ID:        store.ID,
		NamaToko:  store.NamaToko,
		FotoToko:  store.FotoToko,
		CreatedAt: store.CreatedAt,
		UpdatedAt: store.UpdatedAt,
	}
}

// Add FormatStoreResponse function
func FormatStoreResponse(store models.StoreResponse) models.StoreResponse {
	return FormatStore(store)
}
