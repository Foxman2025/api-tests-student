package advertisements

import (
	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// GetAdvertisementsList возвращает список объявлений с учётом поиска и лимита выдачи.
func GetAdvertisementsList(t *testing.T, token, text string, limit int, cursor string) models.AdvertisementsList {
	raw := advertisements.HttpGetAdvertisements(t, token, text, limit, cursor)
	defer raw.Body.Close()

	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err)

	resp := models.AdvertisementsList{}
	require.NoError(t, json.Unmarshal(body, &resp))
	return resp
}

// CreateAdvertisement отправляет запрос на создание объявления и проверяет ожидаемый статус-код
func CreateAdvertisement(t *testing.T, token string, expectedStatusCode int, request models.AdvertisementRequest) string {
	advertisementRawResponse := advertisements.HttpPostAdvertisement(t, token, request)
	defer advertisementRawResponse.Body.Close()

	require.Equalf(t, expectedStatusCode, advertisementRawResponse.StatusCode,
		"HTTP status code должен быть %d", expectedStatusCode)

	respArray, errResponse := io.ReadAll(advertisementRawResponse.Body)
	require.NoError(t, errResponse)

	return string(respArray)
}

// GetAdvertisementByID возвращает объявление по ID и проверяет ожидаемый статус-код
func GetAdvertisementByID(t *testing.T, token string, expectedStatusCode int, id string) string {
	advertisementRawResponse := advertisements.HttpGetAdvertisementById(t, token, id)
	defer advertisementRawResponse.Body.Close()

	require.Equalf(t, expectedStatusCode, advertisementRawResponse.StatusCode,
		"HTTP status code должен быть %d", expectedStatusCode)

	respArray, errResponse := io.ReadAll(advertisementRawResponse.Body)
	require.NoError(t, errResponse)

	return string(respArray)
}

// DeleteAdvertisement удаляет объявление по ID и проверяет ожидаемый статус-код
func DeleteAdvertisement(t *testing.T, token string, expectedStatusCode int, id string) {
	advertisementRawResponse := advertisements.HttpDeleteAdvertisement(t, token, id)
	defer advertisementRawResponse.Body.Close()

	require.Equalf(t, expectedStatusCode, advertisementRawResponse.StatusCode,
		"HTTP status code должен быть %d", expectedStatusCode)
}

// GetAdvertisementPhotos возвращает фотографии объявления и проверяет ожидаемый статус-код
func GetAdvertisementPhotos(t *testing.T, token string, expectedStatusCode int, id string) string {
	photosRawResponse := advertisements.HttpGetAdvertisementPhotosById(t, token, id)
	defer photosRawResponse.Body.Close()

	require.Equalf(t, expectedStatusCode, photosRawResponse.StatusCode,
		"HTTP status code должен быть %d", expectedStatusCode)

	respArray, errResponse := io.ReadAll(photosRawResponse.Body)
	require.NoError(t, errResponse)

	return string(respArray)
}

// DeleteAdvertisementForCleanup удаляет тестовое объявление, если оно ещё существует.
func DeleteAdvertisementForCleanup(t *testing.T, token, id string) {
	t.Helper()

	response := advertisements.HttpDeleteAdvertisement(t, token, id)
	defer response.Body.Close()

	require.Contains(t, []int{http.StatusNoContent, http.StatusNotFound}, response.StatusCode,
		"Не удалось удалить тестовое объявление %s", id)
}
