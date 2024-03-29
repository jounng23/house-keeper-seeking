package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var (
	serverCfg ServerCfg
	dbCfg     DBCfg
)

type DBCfg struct {
	MongoURI string `envconfig:"MONGODB_URI" default:"mongodb://localhost:27017/house_keeper_seeking"`
}

type ServerCfg struct {
	Port          int    `envconfig:"PORT" default:"3000"`
	PricingSvcUrl string `envconfig:"PRICING_SERVICE_URL" default:"http://localhost:3001"`
	SendingSvcUrl string `envconfig:"SENDING_SERVICE_URL" default:"http://localhost:3002"`
}

func InitConfig(configFile string) {
	ReadConfig(configFile)
	configs := []interface{}{
		&serverCfg,
		&dbCfg,
	}
	for _, instance := range configs {
		err := envconfig.Process("", instance)
		if err != nil {
			log.Fatalf("unable to init config: %v, err: %v", instance, err)
		}
	}
}

func ReadConfig(configFile string) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("fatal error config file: %w", err)
	}
	for _, env := range viper.AllKeys() {
		if viper.GetString(env) != "" {
			_ = os.Setenv(env, viper.GetString(env))
			_ = os.Setenv(strings.ToUpper(env), viper.GetString(env))
		}
	}
}

func ServerConfig() ServerCfg {
	return serverCfg
}

func DBConfig() DBCfg {
	return dbCfg
}
