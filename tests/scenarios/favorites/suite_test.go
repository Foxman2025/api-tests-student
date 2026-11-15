package favorites

import (
	"api-tests-template/internal/managers/advertisements"
	advertisementModels "api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/managers/favorites"
	favoriteModels "api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
)

const (
	firstFavoritesPage = 1
	favoritesPageSize  = 50
)

type TestSuite struct {
	suite.Suite
	owner                 authModels.LoginOkResponse
	favoritesUser         authModels.LoginOkResponse
	advertisementID       string
	favoriteAdded         bool
	limitAdvertisementIDs []string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация владельца объявления и пользователя избранного")
	s.owner = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	s.favoritesUser = auth.Login(s.T(), os.Getenv("FAVORITES_TEST_LOGIN"), os.Getenv("FAVORITES_TEST_PASSWORD"))

	require.NotEmpty(s.T(), s.owner.User.Id, "ID владельца пустой")
	require.NotEmpty(s.T(), s.favoritesUser.User.Id, "ID пользователя избранного пустой")
	require.NotEqual(s.T(), s.owner.User.Id, s.favoritesUser.User.Id, "Для теста нужны два разных пользователя")
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) SetupTest() {
	s.advertisementID = ""
	s.favoriteAdded = false
	s.limitAdvertisementIDs = nil
}

// prepareTestAdvertisement создаёт объявление для сценария, которому оно нужно.
func (s *TestSuite) prepareTestAdvertisement() {
	base.Precondition("Создаём объявление от имени владельца")
	s.advertisementID = s.createTestAdvertisement(s.T())
}

func (s *TestSuite) TearDownTest() {
	for _, id := range s.limitAdvertisementIDs {
		defer advertisements.DeleteAdvertisementForCleanup(s.T(), s.owner.Token, id)
	}

	if s.advertisementID == "" {
		return
	}

	s.T().Logf("Убираем данные объявления %s", s.advertisementID)

	defer advertisements.DeleteAdvertisementForCleanup(s.T(), s.owner.Token, s.advertisementID)

	if s.favoriteAdded {
		favorites.DeleteFavorite(s.T(), s.favoritesUser.Token, http.StatusNoContent, s.advertisementID)
	}
}

// addAdvertisementToFavorites добавляет объявление текущего теста в избранное.
func (s *TestSuite) addAdvertisementToFavorites() {
	t := s.T()
	t.Helper()

	request := favoriteModels.AddFavoriteRequest{
		AdvertisementID: s.advertisementID,
	}
	favorites.AddFavorite(t, s.favoritesUser.Token, http.StatusCreated, request)
	s.favoriteAdded = true
}

// getFavoritesBody получает список избранного второго пользователя.
func (s *TestSuite) getFavoritesBody() string {
	return favorites.GetFavorites(
		s.T(), s.favoritesUser.Token, http.StatusOK, "", firstFavoritesPage, favoritesPageSize,
	)
}

// favoriteAdvertisementIDs проверяет массив items и извлекает ID объявлений.
func favoriteAdvertisementIDs(t *testing.T, body string) []string {
	t.Helper()

	items := gjson.Get(body, "items")
	require.True(t, items.IsArray(), "В ответе ожидается массив items")

	var advertisementIDs []string
	for _, item := range items.Array() {
		advertisementIDs = append(advertisementIDs, item.Get("advertisement.id").String())
	}

	return advertisementIDs
}

// createTestAdvertisement создаёт объявление владельца и возвращает его ID.
func (s *TestSuite) createTestAdvertisement(t *testing.T) string {
	t.Helper()

	request := advertisementModels.AdvertisementRequest{
		Title:       "Дом " + utils.RandomString(16),
		Description: "Красивый домик",
		Price:       "555.50",
		Quantity:    "1",
		PhotoPaths:  []string{filepath.Join("..", "advertisementCreate", "testdata", "home1.png")},
	}

	body := advertisements.CreateAdvertisement(t, s.owner.Token, http.StatusCreated, request)
	id := gjson.Get(body, "id").String()
	require.NotEmpty(t, id, "ID созданного объявления пустой")
	t.Logf("Создано объявление %s", id)

	return id
}

// deleteTestAdvertisement удаляет объявление текущего теста и возвращает его ID.
func (s *TestSuite) deleteTestAdvertisement() string {
	t := s.T()
	t.Helper()

	id := s.advertisementID
	advertisements.DeleteAdvertisement(t, s.owner.Token, http.StatusNoContent, id)
	s.advertisementID = ""
	s.favoriteAdded = false

	return id
}
