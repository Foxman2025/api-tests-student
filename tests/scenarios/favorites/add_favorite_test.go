package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	favoriteModels "api-tests-template/internal/managers/favorites/models"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"
)

// Positive: Проверяем, что добавленное объявление появляется в избранном.
func (s *TestSuite) TestAddFavorite() {
	s.prepareTestAdvertisement()
	advertisementID := s.advertisementID

	if !s.Run("Добавляем объявление в избранное", func() {
		s.addAdvertisementToFavorites()
	}) {
		return
	}

	var favoritesBody string
	if !s.Run("Получаем список избранного", func() {
		favoritesBody = s.getFavoritesBody()
	}) {
		return
	}

	s.Run("Проверяем, что добавленное объявление есть в избранном", func() {
		advertisementIDs := favoriteAdvertisementIDs(s.T(), favoritesBody)
		require.Contains(s.T(), advertisementIDs, advertisementID,
			"Добавленное объявление отсутствует в избранном")
	})
}

// Negative: Проверяем отказ при добавлении в избранное с некорректным ID объявления.
func (s *TestSuite) TestAddFavoriteInvalidAdvertisementID() {
	cases := []struct {
		name string
		id   string
	}{
		{name: "пустой ID", id: ""},
		{name: "некорректный UUID", id: "abc"},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			request := favoriteModels.AddFavoriteRequest{AdvertisementID: tc.id}
			expectedMessage := fmt.Sprintf(apierrors.MessageInvalidUUIDLength, len(tc.id))

			var body string
			if !s.Run("Отправляем запрос на добавление в избранное", func() {
				body = favorites.AddFavorite(s.T(), s.favoritesUser.Token, http.StatusBadRequest, request)
			}) {
				return
			}

			s.Run("Проверяем код и сообщение об ошибке", func() {
				assertions.ErrorResponse(s.T(), body, apierrors.CodeInvalidRequest, expectedMessage)
			})
		})
	}
}
