package favorites

import (
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/favorites/models"
	"net/http"
	"strconv"
	"testing"
)

// HttpGetFavorites получает страницу избранного с поиском по названию объявления.
func HttpGetFavorites(t *testing.T, token, textForSearch string, page, pageSize int) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.FavoritesPath).
		Query("search", textForSearch).
		Query("page", strconv.Itoa(page)).
		Query("page_size", strconv.Itoa(pageSize)).
		Expect(t).
		End().Response
}

// HttpPostFavorites добавляет объявление в избранное.
func HttpPostFavorites(t *testing.T, token string, request models.AddFavoriteRequest) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Post(path.FavoritesPath).
		JSON(request).
		Expect(t).
		End().Response
}

// HttpDeleteFavorite удаляет запись из избранного по ID объявления.
func HttpDeleteFavorite(t *testing.T, token, advertisementID string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Deletef(path.FavoriteByAdvertisementIDPath, advertisementID).
		Expect(t).
		End().Response
}

// HttpPostFavoritesBulk добавляет список объявлений в избранное.
func HttpPostFavoritesBulk(t *testing.T, token string, request models.BulkAddFavoritesRequest) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Post(path.FavoritesBulkPath).
		JSON(request).
		Expect(t).
		End().Response
}
