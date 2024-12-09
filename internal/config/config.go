package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	DBConf
	ServiceConf
}

type ServiceConf struct {
	ServicePort string
	GrpcPort    string
}

func NewConfiguration(v *viper.Viper) *Configuration {
	return &Configuration{
		DBConf:      getPostgresDB(v),
		ServiceConf: getConfServer(v),
	}
}

func getConfServer(v *viper.Viper) ServiceConf {

	return ServiceConf{
		ServicePort: v.GetString("server.port"),
		GrpcPort:    v.GetString("server.port-grpc"),
	}

}

func getPostgresDB(v *viper.Viper) DBConf {
	return getConfDBByName("portgres", v)
}
