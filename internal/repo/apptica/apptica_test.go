package apptica

import (
	"context"
	"github.com/shultziv/appposition/internal/domain"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func InSlice(element *domain.AppPosition, s []*domain.AppPosition) bool {
	for _, sEl := range s {
		if reflect.DeepEqual(element, sEl) {
			return true
		}
	}

	return false
}

func TestApptica_GetAppPositionsByDays(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := []byte(`{"status_code":200,"message":"ok","data":{"2":{"1":{"2021-06-01":39}},"23":{"3":{"2021-06-01":51},"1":{"2021-06-01":6}},"134":{"1":{"2021-06-01":72}}}}`)

		w.Write(response)
	}))
	defer ts.Close()

	apiUrl = ts.URL

	var appId uint32 = 1421444
	var countryId uint32 = 1
	dateStr := "2021-06-01"
	date, _ := time.Parse("2006-01-02", dateStr)

	expectedResult := []*domain.AppPosition{
		{
			AppId:         appId,
			CountryId:     countryId,
			CategoryId:    2,
			SubCategoryId: 1,
			Date:          date,
			Position:      39,
		},
		{
			AppId:         appId,
			CountryId:     countryId,
			CategoryId:    23,
			SubCategoryId: 1,
			Date:          date,
			Position:      6,
		},
		{
			AppId:         appId,
			CountryId:     countryId,
			CategoryId:    23,
			SubCategoryId: 3,
			Date:          date,
			Position:      51,
		},
		{
			AppId:         appId,
			CountryId:     countryId,
			CategoryId:    134,
			SubCategoryId: 1,
			Date:          date,
			Position:      72,
		},
	}

	apiKey := "asdasdasdsa"
	apptica := New(apiKey)

	ctx := context.Background()
	result, err := apptica.GetAppPositionsByDays(ctx, appId, countryId, date, date)
	if err != nil {
		t.Error(err)
		return
	}

	for len(expectedResult) != len(result) {
		t.Errorf("Not equal lengths. Expected %d, but result %d", len(expectedResult), len(result))
	}

	for i, _ := range result {
		if !InSlice(result[i], expectedResult) {
			t.Errorf("Invalid result[%d]. Expected %v, but result %v", i, *expectedResult[i], *result[i])
		}
	}
}
