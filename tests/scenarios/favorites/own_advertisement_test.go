package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"net/http"

	"api-tests-template/internal/managers/favorites"
	favoriteModels "api-tests-template/internal/managers/favorites/models"
)

// Negative: Проверяем запрет добавления собственного объявления в избранное.
func (s *TestSuite) TestAddOwnAdvertisementToFavorites() {
	s.prepareTestAdvertisement()
	request := favoriteModels.AddFavoriteRequest{AdvertisementID: s.advertisementID}

	var errorBody string
	if !s.Run("Владелец добавляет своё объявление в избранное", func() {
		errorBody = favorites.AddFavorite(s.T(), s.owner.Token, http.StatusForbidden, request)
	}) {
		return
	}

	s.Run("Проверяем сообщение об ошибке", func() {
		assertions.ErrorResponse(s.T(), errorBody, apierrors.CodeForbidden, apierrors.MessageOwnAdvertisement)
	})
}
