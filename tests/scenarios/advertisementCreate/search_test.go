package advertisementCreate

import (
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"slices"

	"github.com/stretchr/testify/require"
)

// Positive: Проверяем поиск созданного объявления по полному названию
func (s *TestSuite) TestSearchCreatedAdvertisementByFullTitle() {
	title := s.fixture.title
	advertisementID := s.fixture.advertisementResponse.ID

	s.T().Log("Ищем объявление по полному названию")
	advertisementsResponse := advertisements.GetAdvertisementsList(s.T(), s.loginData.Token, title, 0, "")

	s.Run("Среди найденных есть наше объявление", func() {
		found := slices.ContainsFunc(advertisementsResponse.Items, func(item models.Advertisement) bool {
			return item.ID == advertisementID && item.Title == title
		})
		require.Truef(s.T(), found, "Объявление %q с ID %q не найдено", title, advertisementID)
	})
}
