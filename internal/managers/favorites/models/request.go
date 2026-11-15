package models

type AddFavoriteRequest struct {
	AdvertisementID string `json:"advertisement_id"`
}

type BulkAddFavoritesRequest struct {
	AdvertisementIDs []string `json:"advertisement_ids"`
}
