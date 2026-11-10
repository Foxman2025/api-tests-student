package assertions

import (
	"api-tests-template/internal/managers/advertisements/models"
	"testing"

	"github.com/stretchr/testify/require"
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
