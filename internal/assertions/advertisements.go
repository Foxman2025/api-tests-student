package assertions

import (
	"api-tests-template/internal/managers/advertisements/models"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
	t.Helper()

	found := false

	for _, item := range items {
		if item.Title == title {
			found = true
			break
		}
	}

	require.Truef(t, found, "Объявление с названием %q не найдено", title)
}

// ErrorResponse проверяет код и сообщение ошибки в теле ответа
func ErrorResponse(t *testing.T, body, expectedCode, expectedMessage string) {
	t.Helper()

	require.Equal(t, expectedCode, gjson.Get(body, "error").String(),
		"Некорректное значение error параметра")
	require.Equal(t, expectedMessage, gjson.Get(body, "message").String(),
		"Некорректное значение message параметра")
}
