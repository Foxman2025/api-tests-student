package favorites

import (
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"net/http"

	"github.com/stretchr/testify/require"
)

// Positive: Проверяем, что удалённая запись исчезает из избранного.
func (s *TestSuite) TestDeleteFavorite() {
	s.prepareTestAdvertisement()
	advertisementID := s.advertisementID
	base.Precondition("Добавляем объявление в избранное")
	s.addAdvertisementToFavorites()
	beforeDeletion := s.getFavoritesBody()
	require.Contains(s.T(), favoriteAdvertisementIDs(s.T(), beforeDeletion), advertisementID, "объявления нет в избранном")

	var deleteBody string
	if !s.Run("Удаляем объявление из избранного", func() {
		deleteBody = favorites.DeleteFavorite(
			s.T(), s.favoritesUser.Token, http.StatusNoContent, advertisementID,
		)
		s.favoriteAdded = false
	}) {
		return
	}

	var favoritesBody string
	if !s.Run("Получаем список избранного после удаления", func() {
		favoritesBody = s.getFavoritesBody()
	}) {
		return
	}

	s.Run("Проверяем, что объявление исчезло из избранного", func() {
		require.Empty(s.T(), deleteBody, "Тело ответа при 204 должно быть пустым")

		advertisementIDs := favoriteAdvertisementIDs(s.T(), favoritesBody)
		require.NotContains(s.T(), advertisementIDs, advertisementID,
			"Удалённое объявление осталось в избранном")
	})
}

// Positive: Проверяем, что удаление самого объявления убирает его из избранного.
func (s *TestSuite) TestDeletedAdvertisementDisappearsFromFavorites() {
	s.prepareTestAdvertisement()
	advertisementID := s.advertisementID

	base.Precondition("Добавляем чужое объявление в избранное")
	s.addAdvertisementToFavorites()
	beforeDeletion := s.getFavoritesBody()
	require.Contains(s.T(), favoriteAdvertisementIDs(s.T(), beforeDeletion), advertisementID,
		"Предусловие не выполнено: объявления нет в избранном")

	if !s.Run("Удаляем объявление от имени владельца", func() {
		s.deleteTestAdvertisement()
	}) {
		return
	}

	var favoritesBody string
	if !s.Run("Получаем список избранного второго пользователя", func() {
		favoritesBody = s.getFavoritesBody()
	}) {
		return
	}

	s.Run("Проверяем, что удалённое объявление исчезло из избранного", func() {
		advertisementIDs := favoriteAdvertisementIDs(s.T(), favoritesBody)
		require.NotContains(s.T(), advertisementIDs, advertisementID,
			"Удалённое объявление осталось в избранном")
	})
}
