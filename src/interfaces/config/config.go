package config

import (
	"github.com/paraggarg37/order_app/src/interfaces/database"
	"github.com/paraggarg37/order_app/src/interfaces/distancematrix"
	"os"
)

type MainConfig struct {
	Database       database.DBConfig
	DefaultTimeout int
	DistanceMatrix distancematrix.DistanceConfig
}

func GetConfig() *MainConfig {
	return &MainConfig{
		Database: database.DBConfig{
			MasterDSN: os.Getenv("master_db"),
			SlaveDSN:  os.Getenv("slave_db"),
		},
		DefaultTimeout: 30,
		DistanceMatrix: distancematrix.DistanceConfig{
			Url:    "https://maps.googleapis.com",
			ApiKey: os.Getenv("maps_api_key"),
		},
	}
}
