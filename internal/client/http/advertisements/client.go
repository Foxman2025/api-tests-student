package advertisements

import (
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"net/http"
	"testing"
)

//func HttpGetMyAdvertisements(t *testing.T, token string, limit int) *http.Response {
//	return apiRunner.GetRunner().Auth(token).Create().
//		Get(path.MyAdvertisementsPath).
//		Query("limit", strconv.Itoa(limit)).
//		Expect(t).End().Response
//}

func HttpGetAdvertisements(t *testing.T, token, textForSearch string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().
		Get(path.AdvertisementsPath).
		Query("search", textForSearch).
		Expect(t).
		Status(http.StatusOK).
		End().Response
}
