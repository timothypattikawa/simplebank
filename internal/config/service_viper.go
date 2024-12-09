package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadViper() *viper.Viper {

	viperConfig := viper.New()
	viperConfig.SetConfigName("config")
	viperConfig.SetConfigType("yml")
	viperConfig.AddConfigPath(".")
	err := viperConfig.ReadInConfig()
	if err != nil {
		fmt.Printf("Fail to loat viper config yaml err{%v}", err)
		os.Exit(1)
	}

	return viperConfig
}
