package advertisementCreate

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
	"net/http"
)

// Negative: Добавление объявления без авторизации
func (s *TestSuite) TestCreateAdvertisementUnauthorized() {
	cases := []struct {
		name            string
		token           string
		expectedMessage string
	}{
		{"без токена", "", apierrors.MessageAuthHeaderRequired},
		{"неверный токен", "incorrect_token", apierrors.MessageInvalidAuth},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Precondition("Подготовка валидного запроса на создание объявления")
			title := "Дом " + utils.RandomString(8)
			request := s.newAdvertisementRequest(title)

			s.Run("Отправляем запрос и проверяем ошибку авторизации", func() {
				body := advertisements.CreateAdvertisement(s.T(), tc.token, http.StatusUnauthorized, request)

				assertions.ErrorResponse(s.T(), body, apierrors.CodeUnauthorized, tc.expectedMessage)
			})
		})
	}
}
