package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/shultziv/appposition/internal/app"
	"github.com/shultziv/appposition/internal/config"
)

func main() {
	grpcConfig := new(config.GrpcConfig)
	if err := env.Parse(grpcConfig); err != nil {
		fmt.Println(err)
		return
	}

	appticaConfig := new(config.Apptica)
	if err := env.Parse(appticaConfig); err != nil {
		fmt.Println(err)
		return
	}

	postgresConfig := new(config.Postgres)
	if err := env.Parse(postgresConfig); err != nil {
		fmt.Println(err)
		return
	}

	if err := app.Run(grpcConfig, appticaConfig, postgresConfig); err != nil {
		fmt.Println(err)
		return
	}
}
