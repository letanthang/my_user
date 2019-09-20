package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Schema struct {
	PostgresDB struct {
		Username string
		Password string
		Host     string
		Port     int
		Debug    bool
	} `mapstructure:"go_postgres_database"`

	MongoDB struct {
		Username string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"mongo"`

	Paging struct {
		Limit string
	}

	Encryption struct {
		EncryptionKey string
		JWTSecret     string `mapstructure:"jwt_secret"`
	} `mapstructure:"encryption"`
}

var Config Schema

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	// viper.AddConfigPath("/etc/appname/") // path to look for the config file in
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&Config)
}
