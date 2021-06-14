package apptica

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shultziv/appposition/internal/domain"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

var (
	apiUrl  = "https://api.apptica.com"
	timeout = time.Duration(10) * time.Second

	// Колличество дней за которые доступны данные
	countLastDaysWhenDataAvailable = 30
)

// Репозиторий для работы с api apptica
type Apptica struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string) *Apptica {
	return &Apptica{
		apiKey: apiKey,
		httpClient: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   timeout,
		},
	}
}

// Получает данные о рейтингах заданного(appId) приложения в опрделенной стране(countryId) по заданному интервалу времени (дни)
func (a *Apptica) GetAppPositionsByDays(ctx context.Context, appId uint32, countryId uint32, dateFrom time.Time, dateTo time.Time) (appPositions []*domain.AppPosition, err error) {
	validDateBefore := time.Now()
	validDateAfter := validDateBefore.AddDate(0, 0, -countLastDaysWhenDataAvailable)

	// Вмещаются ли даты в достпуный период для API
	if dateTo.Before(validDateAfter) || dateTo.After(validDateBefore) || dateFrom.Before(validDateAfter) || dateFrom.After(validDateBefore) {
		return nil, ErrInvalidDateRange{
			DateFrom: dateFrom,
			DateTo:   dateTo,
		}
	}

	// Вадидны ли даты до и после между собой
	if dateTo.Before(dateFrom) {
		return nil, ErrInvalidDateRange{
			DateFrom: dateFrom,
			DateTo:   dateTo,
		}
	}

	reqData := url.Values{
		"date_from": {dateFrom.Format("2006-01-02")},
		"date_to":   {dateTo.Format("2006-01-02")},
		"B4NKGg":    {a.apiKey},
	}

	topHistoryUrl := fmt.Sprintf("%s/package/top_history", apiUrl)
	reqUrl, err := url.Parse(topHistoryUrl)
	if err != nil {
		return
	}

	reqUrl.Path = path.Join(reqUrl.Path, strconv.Itoa(int(appId)), strconv.Itoa(int(countryId)))
	reqUrl.RawQuery = reqData.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", reqUrl.String(), nil)
	if err != nil {
		return
	}

	response, err := a.httpClient.Do(req)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	topHistoryData := new(topHistoryResponse)
	if err := json.Unmarshal(body, topHistoryData); err != nil {
		return nil, ErrInvalidResponseFromApi{
			ResponseData: body,
		}
	}

	if topHistoryData.StatusCode != 200 {
		return nil, ErrInvalidResponseFromApi{
			ResponseData: body,
		}
	}

	appPositions = make([]*domain.AppPosition, 0)
	for categoryId, subCategoryData := range topHistoryData.Data {
		for subCategoryId, dateData := range subCategoryData {
			for dateStr, position := range dateData {
				date, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					// log
					continue
				}

				appPositions = append(appPositions, &domain.AppPosition{
					AppId:         appId,
					CountryId:     countryId,
					CategoryId:    categoryId,
					SubCategoryId: subCategoryId,
					Date:          date,
					Position:      position,
				})
			}
		}
	}

	return appPositions, nil
}
