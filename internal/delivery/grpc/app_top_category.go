package grpc

import (
	"context"
	pb "github.com/shultziv/appposition/internal/delivery/grpc/proto"
	"google.golang.org/grpc/status"
	"time"
)

const (
	appId     = 1421444
	countryId = 1
)

func (h *Handler) AppTopCategory(ctx context.Context, date *pb.Date) (*pb.CategoryToMaxPosition, error) {
	dateStat, err := time.Parse("2006-01-02", date.Date)
	if err != nil {
		return nil, status.New(402, "Invalid date format.").Err()
	}

	categoryToMaxPosition, err := h.appPositionService.GetMaxPosAppInRangeDays(ctx, appId, countryId, dateStat, dateStat)
	if err != nil {
		return nil, status.New(401, "Couldn't get data").Err()
	}

	return &pb.CategoryToMaxPosition{
		CategoryToMaxPosition: categoryToMaxPosition,
	}, nil
}
