package advertisementCreate

import (
	"api-tests-template/internal/managers/advertisements"
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

// Positive: Проверяем получение созданного объявления по ID
func (s *TestSuite) TestGetCreatedAdvertisementByID() {
	advertisementID := s.fixture.advertisementResponse.ID

	var advertisementBody string
	s.Run("Получаем объявление по ID", func() {
		advertisementBody = advertisements.GetAdvertisementByID(
			s.T(), s.loginData.Token, http.StatusOK, advertisementID,
		)
	})

	s.Run("Проверяем параметры объявления", func() {
		require.Equal(s.T(), advertisementID, gjson.Get(advertisementBody, "id").String(), "id не совпадает")
		require.Equal(s.T(), s.loginData.User.Id, gjson.Get(advertisementBody, "user_id").String(), "user_id не совпадает")
		require.Equal(s.T(), s.fixture.title, gjson.Get(advertisementBody, "title").String(), "Название не совпадает")
		require.Equal(s.T(), advertisementDescription, gjson.Get(advertisementBody, "description").String(), "Описание не совпадает")
		require.Equal(s.T(), advertisementPrice, gjson.Get(advertisementBody, "price").Float(), "Цена не совпадает")
		require.Equal(s.T(), int64(advertisementQuantity), gjson.Get(advertisementBody, "quantity").Int(), "Количество не совпадает")
	})
}
