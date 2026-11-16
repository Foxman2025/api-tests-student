package favorites

import (
	"net/http"
	"slices"
	"time"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	favoriteModels "api-tests-template/internal/managers/favorites/models"

	base "api-tests-template/tests"
)

const (
	maxFavoriteCount             = 50
	limitFixtureInterval         = 250 * time.Millisecond
	messageFavoritesLimitReached = "Favorites limit of 50 reached"
)

// Negative: Проверяем, что API отклоняет добавление 51-го объявления в избранное.
func (s *TestSuite) TestAddFavoriteAboveLimit() {
	t := s.T()
	token := s.favoritesUser.Token
	base.Precondition("Получаем текущее избранное")
	initialIDs := favoriteAdvertisementIDs(t, s.getFavoritesBody())
	require.LessOrEqual(t, len(initialIDs), maxFavoriteCount, "Избранное уже превышает лимит")
	freeSlots := maxFavoriteCount - len(initialIDs)
	base.Precondition("Создаём объявления для заполнения избранного и проверки лимита")
	for i := 0; i < freeSlots+1; i++ {
		if i > 0 {
			time.Sleep(limitFixtureInterval)
		}
		id := s.createTestAdvertisement(t)
		s.limitAdvertisementIDs = append(s.limitAdvertisementIDs, id)
	}
	idsToAdd := s.limitAdvertisementIDs[:freeSlots]
	candidateID := s.limitAdvertisementIDs[freeSlots]
	base.Precondition("Заполняем избранное до 50 записей")
	if len(idsToAdd) > 0 {
		request := favoriteModels.BulkAddFavoritesRequest{AdvertisementIDs: idsToAdd}
		favorites.AddFavoritesBulk(t, token, http.StatusOK, request)
	}
	expectedIDs := slices.Concat(initialIDs, idsToAdd)
	beforeIDs := favoriteAdvertisementIDs(t, s.getFavoritesBody())
	require.ElementsMatch(t, expectedIDs, beforeIDs, "Перед попыткой добавления в избранном должно быть ровно 50 ожидаемых объявлений")

	var errorBody string
	if !s.Run("Пытаемся добавить 51-е объявление в избранное", func() {
		request := favoriteModels.AddFavoriteRequest{AdvertisementID: candidateID}
		errorBody = favorites.AddFavorite(s.T(), token, http.StatusBadRequest, request)
	}) {
		return
	}

	s.Run("Проверяем код и сообщение об ошибке", func() {
		assertions.ErrorResponse(s.T(), errorBody, apierrors.CodeLimitReached, messageFavoritesLimitReached)
	})
	var afterBody string
	if !s.Run("Получаем избранное после попытки добавления", func() {
		afterBody = s.getFavoritesBody()
	}) {
		return
	}
	s.Run("Проверяем, что 51-е объявление не появилось в избранном", func() {
		afterIDs := favoriteAdvertisementIDs(s.T(), afterBody)
		require.ElementsMatch(s.T(), expectedIDs, afterIDs, "В избранном должны остаться те же 50 объявлений")
	})
}
