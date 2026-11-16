package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	favoriteModels "api-tests-template/internal/managers/favorites/models"
	base "api-tests-template/tests"
	"net/http"
	"slices"

	"github.com/stretchr/testify/require"
)

// Negative: Проверяем повторное добавление объявления в избранное.
func (s *TestSuite) TestAddDuplicateFavorite() {
	s.prepareTestAdvertisement()
	advertisementID := s.advertisementID
	request := favoriteModels.AddFavoriteRequest{AdvertisementID: advertisementID}
	base.Precondition("Добавляем чужое объявление в избранное")
	s.addAdvertisementToFavorites()

	var errorBody string
	if !s.Run("Повторно добавляем то же объявление", func() {
		errorBody = favorites.AddFavorite(
			s.T(), s.favoritesUser.Token, http.StatusConflict, request,
		)
	}) {
		return
	}

	var favoritesBody string
	if !s.Run("Получаем список избранного после повторного добавления", func() {
		favoritesBody = s.getFavoritesBody()
	}) {
		return
	}
	s.Run("Проверяем сообщение об ошибке", func() {
		assertions.ErrorResponse(s.T(), errorBody, apierrors.CodeAlreadyExists, apierrors.MessageAlreadyInFavorites)
	})
	s.Run("Проверяем, что объявление осталось в избранном в одном экземпляре", func() {
		advertisementIDs := favoriteAdvertisementIDs(s.T(), favoritesBody)
		matchingIDs := slices.DeleteFunc(advertisementIDs, func(id string) bool {
			return id != advertisementID
		})
		require.Len(s.T(), matchingIDs, 1, "Объявление должно быть в избранном ровно один раз")
	})
}
