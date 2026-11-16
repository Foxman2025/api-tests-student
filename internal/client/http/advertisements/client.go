package advertisements

import (
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/advertisements/models"
	"net/http"
	"strconv"
	"testing"
)

// HttpGetAdvertisements получает страницу объявлений с поиском и лимитом.
func HttpGetAdvertisements(t *testing.T, token, textForSearch string, limit int, cursor string) *http.Response {
	req := apiRunner.GetRunner().Auth(token).Create().Get(path.AdvertisementsPath)
	req.Query("search", textForSearch)
	if limit != 0 {
		req.Query("limit", strconv.Itoa(limit))
	}
	if cursor != "" {
		req.Query("cursor", cursor)
	}

	return req.Expect(t).Status(http.StatusOK).End().Response
}

// HttpPostAdvertisement отправляет multipart-запрос на создание объявления.
func HttpPostAdvertisement(t *testing.T, token string, request models.AdvertisementRequest) *http.Response {
	req := apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath)

	for name, value := range request.FormFields() {
		req.MultipartFormData(name, value)
	}

	req.MultipartFile(models.FormFieldPhotos, request.PhotoPaths...)

	return req.Expect(t).End().Response
}

// HttpGetAdvertisementById получает объявление по id.
func HttpGetAdvertisementById(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

// HttpDeleteAdvertisement удаляет объявление по id.
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Delete(path.AdvertisementPath).
		Query("id", id).
		Expect(t).
		End().Response
}

// HttpGetAdvertisementPhotosById получает фотографии объявления по его id.
func HttpGetAdvertisementPhotosById(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Getf(path.AdvertisementPhotosPath, id).
		Expect(t).
		End().Response
}
