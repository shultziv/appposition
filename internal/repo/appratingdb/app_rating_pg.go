package appratingdb

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shultziv/appposition/internal/domain"
	"time"
)

type AppRatingPg struct {
	pg *pgxpool.Pool
}

func NewAppRatingPg(pg *pgxpool.Pool) *AppRatingPg {
	return &AppRatingPg{
		pg: pg,
	}
}

func (a *AppRatingPg) GetMaxPosAppByDays(ctx context.Context, appId uint32, countryId uint32, dateFrom time.Time, dateTo time.Time) (categoryIdToMaxPos map[uint32]uint32, err error) {
	rows, err := a.pg.Query(ctx,
		`SELECT 
				category_id, 
				MIN(pos) 
			FROM 
				app_rating 
			WHERE 
				date_rate 
				BETWEEN 
					$3 
					AND 
					$4 
				AND
					app_id = $1 
				AND 
					country_id = $2
			GROUP BY 
				app_id, 
				country_id, 
				category_id;`, appId, countryId, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categoryIdToMaxPos = make(map[uint32]uint32)
	for rows.Next() {
		var categoryId uint32
		var maxPos uint32
		if err = rows.Scan(&categoryId, &maxPos); err != nil {
			return nil, err
		}
		categoryIdToMaxPos[categoryId] = maxPos
	}

	if len(categoryIdToMaxPos) == 0 {
		return nil, ErrDataNotFound
	}

	return categoryIdToMaxPos, nil
}

func (a *AppRatingPg) AddAppPositions(ctx context.Context, appPositions []*domain.AppPosition) (err error) {
	rows := make([][]interface{}, 0)
	columnNames := []string{"app_id", "country_id", "category_id", "sub_category_id", "date_rate", "pos"}
	for _, appPosition := range appPositions {
		rows = append(rows, []interface{}{
			appPosition.AppId,
			appPosition.CountryId,
			appPosition.CategoryId,
			appPosition.SubCategoryId,
			appPosition.Date,
			appPosition.Position,
		})
	}

	_, err = a.pg.CopyFrom(ctx, pgx.Identifier{"app_rating"}, columnNames, pgx.CopyFromRows(rows))
	if err != nil {
		return err
	}

	return nil
}
