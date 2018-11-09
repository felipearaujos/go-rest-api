package configuration

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	viper.SetConfigFile("Config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	mongo := viper.Sub("mongo")
	mongo.Unmarshal(&c)
}
