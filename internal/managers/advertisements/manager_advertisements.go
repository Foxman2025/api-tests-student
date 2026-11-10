package advertisements

import (
	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

//func Login(t *testing.T, username, password string) models.LoginOkResponse {
//	raw := auth.HttpPostAuthLogin(t, models.LoginRequest{Email: username, Password: password})
//	defer raw.Body.Close()
//	body, err := io.ReadAll(raw.Body); require.NoError(t, err)
//	resp := models.LoginOkResponse{}
//	require.NoError(t, json.Unmarshal(body, &resp))
//	return resp
//}

func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	raw := advertisements.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err)
	resp := models.AdvertisementsList{}
	require.NoError(t, json.Unmarshal(body, &resp))
	return resp
}
