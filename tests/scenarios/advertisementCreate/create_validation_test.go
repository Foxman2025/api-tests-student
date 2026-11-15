package advertisementCreate

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/apierrors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/utils"

	base "api-tests-template/tests"
)

const (
	maxTitleLength             = 50
	maxDescriptionLength       = 500
	messageTitleRequired       = "title is required"
	messageDescriptionRequired = "description is required"
	messagePhotoRequired       = "at least 1 photo is required"
	messageTitleTooLongFormat  = "title must not exceed %d characters provided %d"
	messageDescriptionLength   = "description must be between 1 and 500 characters"
	messagePriceRange          = "price must be between 0 and 1000000 inclusive"
	messageQuantityRange       = "quantity must be between 1 and 100 inclusive"
	messagePhotosMaxCount      = "maximum 3 photos allowed"
	messagePriceNotNumber      = "cannot convert string value abc to float"
	messageFileTypeNotAllowed  = `photo 1: file type not allowed: extension ".txt"`
)

// Negative: Проверяем ошибки валидации при создании объявления.
func (s *TestSuite) TestCreateAdvertisementInvalidData() {
	testCases := []struct {
		name            string
		expectedMessage string
		changeRequest   func(*testing.T, *models.AdvertisementRequest)
	}{
		{
			name:            "пустое название",
			expectedMessage: messageTitleRequired,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.Title = ""
			},
		},
		{
			name:            "название отсутствует",
			expectedMessage: messageTitleRequired,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.OmitFields = []string{models.FormFieldTitle}
			},
		},
		{
			name:            "описание отсутствует",
			expectedMessage: messageDescriptionRequired,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.OmitFields = []string{models.FormFieldDescription}
			},
		},
		{
			name:            "название длиннее 50 символов",
			expectedMessage: fmt.Sprintf(messageTitleTooLongFormat, maxTitleLength, maxTitleLength+1),
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.Title = strings.Repeat("a", maxTitleLength+1)
			},
		},
		{
			name:            "описание длиннее 500 символов",
			expectedMessage: messageDescriptionLength,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.Description = strings.Repeat("a", maxDescriptionLength+1)
			},
		},
		{
			name:            "отрицательная цена",
			expectedMessage: messagePriceRange,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.Price = "-1"
			},
		},
		{
			name:            "цена передана текстом",
			expectedMessage: messagePriceNotNumber,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.Price = "abc"
			},
		},
		{
			name:            "количество равно нулю",
			expectedMessage: messageQuantityRange,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.Quantity = "0"
			},
		},
		{
			name:            "фотографии отсутствуют",
			expectedMessage: messagePhotoRequired,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.PhotoPaths = nil
			},
		},
		{
			name:            "четыре фотографии",
			expectedMessage: messagePhotosMaxCount,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.PhotoPaths = []string{
					filepath.Join("testdata", "home1.png"),
					filepath.Join("testdata", "home2.png"),
					filepath.Join("testdata", "home3.png"),
					filepath.Join("testdata", "home4.png"),
				}
			},
		},
		{
			name:            "фотография с расширением txt",
			expectedMessage: messageFileTypeNotAllowed,
			changeRequest: func(_ *testing.T, request *models.AdvertisementRequest) {
				request.PhotoPaths = []string{filepath.Join("testdata", "wrong_extension.txt")}
			},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			t := s.T()

			base.Precondition("Подготовка запроса с некорректными данными")
			request := s.newAdvertisementRequest("Дом " + utils.RandomString(8))
			testCase.changeRequest(t, &request)

			body := advertisements.CreateAdvertisement(t, s.loginData.Token, http.StatusBadRequest, request)

			t.Logf("Ответ API: %s", body)
			assertions.ErrorResponse(t, body, apierrors.CodeInvalidRequest, testCase.expectedMessage)
		})
	}
}
