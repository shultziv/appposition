package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shultziv/appposition/internal/config"
	"github.com/shultziv/appposition/internal/delivery/grpc"
	"github.com/shultziv/appposition/internal/repo/appratingdb"
	"github.com/shultziv/appposition/internal/repo/apptica"
	"github.com/shultziv/appposition/internal/server"
	"github.com/shultziv/appposition/internal/service/appposition"
)

func Run(grpcConfig *config.GrpcConfig, appticaConfig *config.Apptica, postgresConfig *config.Postgres) (err error) {
	appRatingRepo := apptica.New(appticaConfig.ApiKey)

	pg, err := pgxpool.Connect(context.Background(), postgresConfig.Url)
	if err != nil {
		return
	}

	appRatingDbRepo := appratingdb.NewAppRatingPg(pg)

	appPositionService := appposition.New(appRatingRepo, appRatingDbRepo)

	grpcHandler := grpc.NewHandler(appPositionService)

	grpcServer := server.NewGrpcServer(grpcConfig.Port)
	grpcHandler.RegisterServer(grpcServer.Server)

	return grpcServer.Run()
}
