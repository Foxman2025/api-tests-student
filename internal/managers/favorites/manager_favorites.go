package favorites

import (
	"api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/managers/favorites/models"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// GetFavorites получает страницу избранного и проверяет ожидаемый статус-код.
func GetFavorites(t *testing.T, token string, expectedStatusCode int, textForSearch string, page, pageSize int) string {
	favoritesRawResponse := favorites.HttpGetFavorites(t, token, textForSearch, page, pageSize)
	return readFavoritesResponse(t, favoritesRawResponse, expectedStatusCode)
}

// AddFavorite добавляет объявление в избранное и проверяет ожидаемый статус-код.
func AddFavorite(t *testing.T, token string, expectedStatusCode int, request models.AddFavoriteRequest) string {
	favoritesRawResponse := favorites.HttpPostFavorites(t, token, request)
	return readFavoritesResponse(t, favoritesRawResponse, expectedStatusCode)
}

// DeleteFavorite удаляет объявление из избранного и проверяет ожидаемый статус-код.
func DeleteFavorite(t *testing.T, token string, expectedStatusCode int, advertisementID string) string {
	favoritesRawResponse := favorites.HttpDeleteFavorite(t, token, advertisementID)
	return readFavoritesResponse(t, favoritesRawResponse, expectedStatusCode)
}

// AddFavoritesBulk добавляет несколько объявлений в избранное и проверяет ожидаемый статус-код.
func AddFavoritesBulk(t *testing.T, token string, expectedStatusCode int, request models.BulkAddFavoritesRequest) string {
	favoritesRawResponse := favorites.HttpPostFavoritesBulk(t, token, request)
	return readFavoritesResponse(t, favoritesRawResponse, expectedStatusCode)
}

// readFavoritesResponse проверяет статус-код, читает и закрывает тело ответа.
func readFavoritesResponse(t *testing.T, response *http.Response, expectedStatusCode int) string {
	t.Helper()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	require.Equalf(t, expectedStatusCode, response.StatusCode,
		"HTTP status code должен быть %d. Ответ API: %s", expectedStatusCode, body)

	return string(body)
}
