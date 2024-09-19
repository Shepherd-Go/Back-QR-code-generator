package config

import (
	"log"
	"os"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

type Config struct {
	Server  Server  `required:"true" mapstructure:"server"`
	Databse Databse `required:"true" mapstructure:"database"`
}

type Server struct {
	Port int `required:"true" mapstructure:"port"`
}

type Databse struct {
	MongoDBConnectionWrite string `required:"true" mapstructure:"mongo_db_connection_write"`
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {

		if os.Getenv("DEVMODE") == "true" {
			if err := config.SetEnvsFromFile("qr-system", "../.env"); err != nil {
				log.Panic(err)
			}
		}

		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panicf("Error parsing enviroment vars %v", err)
		}
	})

	return Cfg
}
