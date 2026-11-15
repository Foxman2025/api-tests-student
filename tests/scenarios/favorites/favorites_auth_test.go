package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/favorites"
	favoriteModels "api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/utils"
	"net/http"
	"testing"
)

// Negative: Проверяем, что ручки избранного требуют авторизацию.
func (s *TestSuite) TestFavoritesWithoutAuthorization() {
	id := utils.RandomUUID()
	const token = ""

	cases := []struct {
		name string
		send func(*testing.T) string
	}{
		{"GET /favorites", func(t *testing.T) string {
			return favorites.GetFavorites(
				t, token, http.StatusUnauthorized, "", firstFavoritesPage, favoritesPageSize,
			)
		}},
		{"POST /favorites", func(t *testing.T) string {
			return favorites.AddFavorite(t, token, http.StatusUnauthorized,
				favoriteModels.AddFavoriteRequest{AdvertisementID: id})
		}},
		{"DELETE /favorites/{ad_id}", func(t *testing.T) string {
			return favorites.DeleteFavorite(t, token, http.StatusUnauthorized, id)
		}},
		{"POST /favorites/bulk", func(t *testing.T) string {
			return favorites.AddFavoritesBulk(t, token, http.StatusUnauthorized,
				favoriteModels.BulkAddFavoritesRequest{AdvertisementIDs: []string{id}})
		}},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			var body string
			if !s.Run("Отправляем запрос без токена", func() {
				body = tc.send(s.T())
			}) {
				return
			}

			s.Run("Проверяем код и сообщение об ошибке", func() {
				assertions.ErrorResponse(s.T(), body, apierrors.CodeUnauthorized, apierrors.MessageAuthHeaderRequired)
			})
		})
	}
}
