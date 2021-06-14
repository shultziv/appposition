package appposition

import (
	"context"
	"github.com/shultziv/appposition/internal/domain"
	"time"
)

//go:generate mockgen -destination=mocks/mock_app_rating_repo.go -package=mocks . AppRatingRepo
type AppRatingRepo interface {
	GetAppPositionsByDays(ctx context.Context, appId uint32, countryId uint32, dateFrom time.Time, dateTo time.Time) (appPositions []*domain.AppPosition, err error)
}

//go:generate mockgen -destination=mocks/mock_app_rating_db_repo.go -package=mocks . AppRatingDbRepo
type AppRatingDbRepo interface {
	GetMaxPosAppByDays(ctx context.Context, appId uint32, countryId uint32, dateFrom time.Time, dateTo time.Time) (categoryIdToMaxPos map[uint32]uint32, err error)
	AddAppPositions(ctx context.Context, appPositions []*domain.AppPosition) (err error)
}

type AppPosition struct {
	appRatingRepo   AppRatingRepo
	appRatingDbRepo AppRatingDbRepo
}

func New(appRatingRepo AppRatingRepo, appRatingDbRepo AppRatingDbRepo) *AppPosition {
	return &AppPosition{
		appRatingRepo:   appRatingRepo,
		appRatingDbRepo: appRatingDbRepo,
	}
}

func calcMaxPosApp(appPositions []*domain.AppPosition) (categoryIdToMaxPos map[uint32]uint32) {
	categoryIdToMaxPos = make(map[uint32]uint32)
	for _, appPosition := range appPositions {
		if appPosition.Position < categoryIdToMaxPos[appPosition.CategoryId] || categoryIdToMaxPos[appPosition.CategoryId] == 0 {
			categoryIdToMaxPos[appPosition.CategoryId] = appPosition.Position
		}
	}

	return categoryIdToMaxPos
}

func (ap *AppPosition) GetMaxPosAppInRangeDays(ctx context.Context, appId uint32, countryId uint32, dateFrom time.Time, dateTo time.Time) (categoryIdToMaxPos map[uint32]uint32, err error) {
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

	// Кладем в БД ответ через горутину, чтобы не задерживать ответ
	go func(ctx context.Context, appPositions []*domain.AppPosition) {
		if err := ap.appRatingDbRepo.AddAppPositions(ctx, appPositions); err != nil {
			// Если даные не сохранятся в БД - не критично, самое главное нам как-то об этом узнать
			// Просто залогируем
			// TODO: log
		}
	}(context.Background(), appPositions)

	return calcMaxPosApp(appPositions), nil
}
