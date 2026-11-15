package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	favoriteModels "api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/utils"
	"net/http"
)

// Negative: Проверяем добавление несуществующего объявления в избранное.
func (s *TestSuite) TestAddMissingAdvertisementToFavorites() {
	id := utils.RandomUUID()
	request := favoriteModels.AddFavoriteRequest{AdvertisementID: id}

	var body string
	if !s.Run("Добавляем несуществующее объявление в избранное", func() {
		body = favorites.AddFavorite(s.T(), s.favoritesUser.Token, http.StatusNotFound, request)
	}) {
		return
	}

	s.Run("Проверяем сообщение об ошибке", func() {
		assertions.ErrorResponse(s.T(), body, apierrors.CodeNotFound, apierrors.MessageAdvertisementNotFound)
	})
}

// Negative: Проверяем удаление несуществующего объявления из избранного.
func (s *TestSuite) TestDeleteMissingAdvertisementFromFavorites() {
	id := utils.RandomUUID()

	var body string
	if !s.Run("Удаляем запись избранного по ID несуществующего объявления", func() {
		body = favorites.DeleteFavorite(s.T(), s.favoritesUser.Token, http.StatusNotFound, id)
	}) {
		return
	}

	s.Run("Проверяем сообщение об ошибке", func() {
		assertions.ErrorResponse(s.T(), body, apierrors.CodeNotFound, apierrors.MessageFavoriteNotFound)
	})
}
