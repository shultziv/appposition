package appposition

import (
	"context"
	"github.com/shultziv/appposition/internal/domain"
	"time"
)

type appRatingRepo interface {
	GetAppPositionsByDays(ctx context.Context, appId uint32, countryId uint8, dateFrom time.Time, dateTo time.Time) (appPositions []*domain.AppPosition, err error)
}

type appRatingDbRepo interface {
	GetMaxPosAppByDays(ctx context.Context, appId uint32, countryId uint8, dateFrom time.Time, dateTo time.Time) (categoryIdToMaxPos map[uint8]uint8, err error)
	AddAppPositions(ctx context.Context, appPositions []*domain.AppPosition) (err error)
}

type AppPosition struct {
	appRatingRepo   appRatingRepo
	appRatingDbRepo appRatingDbRepo
}

func New(appRatingRepo appRatingRepo, appRatingDbRepo appRatingDbRepo) *AppPosition {
	return &AppPosition{
		appRatingRepo:   appRatingRepo,
		appRatingDbRepo: appRatingDbRepo,
	}
}

func (ap *AppPosition) GetMaxPosAppInRecentDays(ctx context.Context, appId uint32, countryId uint8, countRecentDays uint32) (categoryIdToMaxPos map[uint8]uint8, err error) {
	dateTo := time.Now()
	dateFrom := dateTo.AddDate(0, 0, int(-countRecentDays))

	// Сходили в БД, если все ок (данные есть за нужные дни), то отдаем результат
	categoryIdToMaxPos, err = ap.appRatingDbRepo.GetMaxPosAppByDays(ctx, appId, countryId, dateFrom, dateTo)
	if err == nil {
		return categoryIdToMaxPos, err
	}

	// Если данных в бд нет или недостаточно, то идем в "активный" репозиторий
	appPositions, err := ap.appRatingRepo.GetAppPositionsByDays(ctx, appId, countryId, dateFrom, dateTo)
	if err != nil {
		// TODO: log
		return nil, ErrCouldNotGetData
	}

	go func(ctx context.Context, appPositions []*domain.AppPosition) {
		if err := ap.appRatingDbRepo.AddAppPositions(ctx, appPositions); err != nil {
			// Если даные не сохранятся в БД - не критично, самое главное нам как-то об этом узнать
			// Просто залогируем
			// TODO: log
		}
	}(ctx, appPositions)

	return
}
