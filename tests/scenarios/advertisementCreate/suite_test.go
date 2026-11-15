package advertisementCreate

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"

	base "api-tests-template/tests"
)

const (
	advertisementDescription = "Красивый домик"
	advertisementPrice       = 555.50
	advertisementQuantity    = 1
)

type TestSuite struct {
	suite.Suite
	loginData authModels.LoginOkResponse
	fixture   advertisementFixture
}

type advertisementFixture struct {
	title                 string
	advertisementResponse models.Advertisement
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация пользователя с кредами из переменных окружения и получение его параметров")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	s.fixture = s.createAdvertisementFixture()
}

func (s *TestSuite) TearDownSuite() {
	defer base.TearDownSuite()

	advertisementID := s.fixture.advertisementResponse.ID
	s.T().Logf("Удаляем объявление %s", advertisementID)
	advertisements.DeleteAdvertisement(s.T(), s.loginData.Token, http.StatusNoContent, advertisementID)
}

// createAdvertisementFixture создаёт объявление со всеми параметрами и фотографией
func (s *TestSuite) createAdvertisementFixture() advertisementFixture {
	t := s.T()
	t.Helper()

	base.Precondition("Создаём объявление со всеми параметрами и фотографией")

	title := "Дом " + utils.RandomString(8)
	request := s.newAdvertisementRequest(title)

	advertisementBody := advertisements.CreateAdvertisement(
		t, s.loginData.Token, http.StatusCreated, request,
	)

	advertisementResponse := models.Advertisement{}
	err := json.Unmarshal([]byte(advertisementBody), &advertisementResponse)
	require.NoError(t, err)
	require.NotEmpty(t, advertisementResponse.ID, "ID созданного объявления пустой")

	t.Logf("Создано объявление %s", advertisementResponse.ID)

	return advertisementFixture{
		title:                 title,
		advertisementResponse: advertisementResponse,
	}
}

// newAdvertisementRequest подготавливает валидный запрос на создание объявления
func (s *TestSuite) newAdvertisementRequest(title string) models.AdvertisementRequest {
	return models.AdvertisementRequest{
		Title:       title,
		Description: advertisementDescription,
		Price:       strconv.FormatFloat(advertisementPrice, 'f', -1, 64),
		Quantity:    strconv.Itoa(advertisementQuantity),
		PhotoPaths:  []string{filepath.Join("testdata", "home1.png")},
	}
}
