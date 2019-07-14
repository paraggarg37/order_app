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
			MasterDSN: os.Getenv("MASTER_DB"),
			SlaveDSN:  os.Getenv("SLAVE_DB"),
		},
		DefaultTimeout: 30,
		DistanceMatrix: distancematrix.DistanceConfig{
			Url:    "https://maps.googleapis.com",
			ApiKey: os.Getenv("MAPS_API_KEY"),
		},
	}
}
