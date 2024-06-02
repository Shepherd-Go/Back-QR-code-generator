package config

import (
	"sync"

	"github.com/andresxlp/gosuite/config"
)

type Config struct {
	ServerHost             string `required:"true" mapstructure:"server_host"`
	Port                   int    `required:"true" mapstructure:"port"`
	MongoDBConnectionWrite string `required:"true" mapstructure:"mongo_db_connection_write"`
	InternalPrivatePath    string `required:"true" mapstructure:"internal_private_path"`
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {
		config.SetEnvsFromFile("qr-system", ".env")
		config.GetConfigFromEnv(&Cfg)
	})

	return Cfg
}
