// @title Logistics API
// @version 1.0
// @description Logistics Service API
// @host localhost:8080
// @BasePath /
//go:generate sqlboiler --config  ./../sqlboiler.toml --output ./../src/db_modles --pkgname db_modles  --wipe  mysql --no-tests

package main

import (
	"github.com/paraggarg37/order_app/src/domain/repositories"
	"github.com/paraggarg37/order_app/src/interfaces/config"
	"github.com/paraggarg37/order_app/src/interfaces/dao"
	"github.com/paraggarg37/order_app/src/interfaces/database"
	"github.com/paraggarg37/order_app/src/interfaces/distancematrix"
	"github.com/paraggarg37/order_app/src/interfaces/web/api"
	"github.com/paraggarg37/order_app/src/usecases/order"
	"log"
)

func main() {
	oApi := Setup()
	oApi.RegisterOrder()
	oApi.Run()
}

func Setup() *api.API {
	cfg := config.GetConfig()

	db, err := database.New(cfg.Database, database.DriverMySQL)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	orderDao := dao.New(db)

	dService, err := distancematrix.Init(cfg.DistanceMatrix.Url, cfg.DistanceMatrix.ApiKey)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	oApi := getAPI(cfg, dService, orderDao, db)

	return oApi
}

func getAPI(cfg *config.MainConfig, distanceService distancematrix.DistanceService, repo repositories.OrderRepository, db database.DBRepository) *api.API {
	return api.New(&api.API{
		Cfg: cfg,
		Interactor: &api.Interactor{
			OrderInteractor: order.Init(distanceService, repo),
		},
		DB: db,
	})
}
