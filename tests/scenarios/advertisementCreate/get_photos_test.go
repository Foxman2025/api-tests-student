package advertisementCreate

import (
	"api-tests-template/internal/managers/advertisements"
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

// Positive: Проверяем получение фотографии созданного объявления
func (s *TestSuite) TestGetCreatedAdvertisementPhotos() {
	advertisementID := s.fixture.advertisementResponse.ID
	uploadedPhotos := s.fixture.advertisementResponse.Photos
	require.Len(s.T(), uploadedPhotos, 1, "В ответе на создание должна быть одна фотография")
	expectedPhotoID := uploadedPhotos[0].ID
	require.NotEmpty(s.T(), expectedPhotoID, "ID загруженной фотографии пустой")

	var photosBody string
	if !s.Run("Получаем список фотографий объявления", func() {
		photosBody = advertisements.GetAdvertisementPhotos(
			s.T(), s.loginData.Token, http.StatusOK, advertisementID,
		)
	}) {
		return
	}

	s.Run("Проверяем, что в списке есть загруженная фотография", func() {
		photos := gjson.Parse(photosBody)
		require.True(s.T(), photos.IsArray(), "В ответе ожидается массив фотографий")

		items := photos.Array()
		require.Len(s.T(), items, 1, "В списке должна быть одна фотография")
		require.Equal(s.T(), expectedPhotoID, items[0].Get("id").String(), "ID фотографии не совпадает")
		require.NotEmpty(s.T(), items[0].Get("url").String(), "URL фотографии пустой")
	})
}
