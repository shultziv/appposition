package appposition

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shultziv/appposition/internal/domain"
	"github.com/shultziv/appposition/internal/service/appposition/mocks"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestAppPosition_GetMaxPosAppInRecentDaysExistInDb(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppRatingDbRepo := mocks.NewMockAppRatingDbRepo(ctrl)
	mockAppRatingRepo := mocks.NewMockAppRatingRepo(ctrl)

	appPositionService := New(mockAppRatingRepo, mockAppRatingDbRepo)

	var appId uint32 = 177
	var countryId uint32 = 1
	ctx := context.Background()
	expectedResult := map[uint32]uint32{
		1: 1,
		2: 2,
	}

	mockAppRatingDbRepo.
		EXPECT().
		GetMaxPosAppByDays(gomock.Any(), gomock.Eq(appId), gomock.Eq(countryId), gomock.Any(), gomock.Any()).
		Return(expectedResult, nil)

	result, err := appPositionService.GetMaxPosAppInRangeDays(ctx, appId, countryId, time.Now().AddDate(0, 0, -1), time.Now())
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Result %v not equal expected %v", result, expectedResult)
	}
}

func TestAppPosition_GetMaxPosAppInRecentDaysNotExistInDb(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup
	defer wg.Wait()

	mockAppRatingDbRepo := mocks.NewMockAppRatingDbRepo(ctrl)
	mockAppRatingRepo := mocks.NewMockAppRatingRepo(ctrl)

	appPositionService := New(mockAppRatingRepo, mockAppRatingDbRepo)

	var appId uint32 = 177
	var countryId uint32 = 1
	ctx := context.Background()

	appPositions := []*domain.AppPosition{
		{
			AppId:         appId,
			CountryId:     countryId,
			CategoryId:    1,
			SubCategoryId: 1,
			Date:          time.Now(),
			Position:      1,
		},
		{
			AppId:         appId,
			CountryId:     countryId,
			CategoryId:    1,
			SubCategoryId: 2,
			Date:          time.Now(),
			Position:      4,
		},
		{
			AppId:         appId,
			CountryId:     countryId,
			CategoryId:    4,
			SubCategoryId: 2,
			Date:          time.Now(),
			Position:      3,
		},
	}

	expectedResult := map[uint32]uint32{
		1: 1,
		4: 3,
	}

	mockAppRatingDbRepo.
		EXPECT().
		GetMaxPosAppByDays(gomock.Any(), gomock.Eq(appId), gomock.Eq(countryId), gomock.Any(), gomock.Any()).
		Return(nil, errors.New("Any error"))

	wg.Add(1)
	mockAppRatingDbRepo.
		EXPECT().
		AddAppPositions(gomock.Any(), gomock.Any()).
		Do(func(ctx context.Context, appPositions []*domain.AppPosition) {
			wg.Done()
		}).
		Return(nil)

	mockAppRatingRepo.
		EXPECT().
		GetAppPositionsByDays(ctx, appId, countryId, gomock.Any(), gomock.Any()).
		Return(appPositions, nil)

	result, err := appPositionService.GetMaxPosAppInRangeDays(ctx, appId, countryId, time.Now().AddDate(0, 0, -1), time.Now())
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Result %v not equal expected %v", result, expectedResult)
	}
}
